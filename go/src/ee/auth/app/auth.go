package app

import (
	"ee/auth"
	"errors"
	"github.com/go-ee/utils/crypt"
	"github.com/go-ee/utils/eh/app"
	"github.com/go-ee/utils/net"
	"github.com/google/uuid"
)

type Auth struct {
	*app.AppBase
}

func NewAuth(appBase *app.AppBase) *Auth {
	appBase.ProductName = "Auth"
	return &Auth{AppBase: appBase}
}

func (o *Auth) Start() (err error) {

	authEngine := auth.NewEsInitializer(o.EventStore, o.EventBus, o.CommandBus, o.ReadRepos)
	if err = authEngine.Setup(); err != nil {
		return
	}

	authEngine.ActivatePasswordEncryption()

	authRouter := auth.NewRouter("", o.NewContext, o.CommandBus, o.ReadRepos)
	if err = authRouter.Setup(o.Router); err != nil {
		return
	}

	if o.Secure {
		o.Jwt = o.initJwtController(authRouter.AccountRouter.QueryHandler.QueryRepository)
	}

	err = o.StartServer()
	return
}

func (o *Auth) initJwtController(accounts *auth.AccountQueryRepository) (ret *net.JwtController) {
	//TODO use appName, provide help how to generate RSA files first
	return net.NewJwtControllerApp("app",
		func(credentials net.UserCredentials) (ret interface{}, err error) {
			var account *auth.Account
			id, _ := uuid.Parse(credentials.Username)
			if account, err = accounts.FindById(id); err == nil {
				if !crypt.HashAndEquals(credentials.Password, account.Password) {
					err = errors.New("password mismatch")
				} else {
					ret = account
				}
			}
			return
		})
}
