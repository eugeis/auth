package app

import (
	"ee/auth"
	"errors"
	"github.com/go-ee/utils/crypt"
	"github.com/go-ee/utils/eh/app"
	"github.com/go-ee/utils/net"
	"github.com/google/uuid"
	"path/filepath"
)

type Auth struct {
	*app.AppBase
}

func NewAuth(appBase *app.AppBase) *Auth {
	appBase.ProductName = "Auth"
	return &Auth{AppBase: appBase}
}

func (o *Auth) Start() (err error) {

	authEngine := auth.NewEsEngine(o.Middleware)
	if err = authEngine.Setup(); err != nil {
		return
	}

	authEngine.ActivatePasswordEncryption()
	var authRouter *auth.Router
	if authRouter, err = auth.NewRouter("", o.NewContext, authEngine); err != nil {
		return
	}
	if err = authRouter.Setup(o.Router); err != nil {
		return
	}

	if o.Secure {
		if o.Jwt, err = o.initJwtController(authRouter.AccountRouter.QueryHandler.QueryRepository); err != nil {
			return
		}
	}

	err = o.StartServer()
	return
}

func (o *Auth) initJwtController(accounts *auth.AccountQueryRepository) (ret *net.JwtController, err error) {
	return net.NewJwtControllerApp(
		filepath.Join(o.WorkingFolder, "certs"), o.AppName,
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
