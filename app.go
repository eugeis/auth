package main

import (
	"errors"

	"github.com/eugeis/auth/auth"
	"github.com/eugeis/gee/crypt"
	"github.com/eugeis/gee/eh/app"
	"github.com/eugeis/gee/net"
	"github.com/looplab/eventhorizon"
)

type Auth struct {
	*app.AppBase
}

func NewAuth(appBase *app.AppBase) *Schkola {
	appBase.ProductName = "Auth"
	return &Auth{AppBase: appBase}
}

func (o *Auth) Start() {

	authEngine := auth.NewAuthEventhorizonInitializer(o.EventStore, o.EventBus, o.CommandBus, o.ReadRepos)
	authEngine.Setup()
	authEngine.ActivatePasswordEncryption()

	if o.Secure {
		o.Jwt = o.initJwtController(authRouter.AccountRouter.QueryHandler.QueryRepository)
	}

	o.StartServer()
}

func (o *Auth) initJwtController(accounts *auth.AccountQueryRepository) (ret *net.JwtController) {
	//TODO use appName, provide help how to generate RSA files first
	return net.NewJwtControllerApp("app",
		func(credentials net.UserCredentials) (ret interface{}, err error) {
			var account *auth.Account
			if account, err = accounts.FindById(eventhorizon.UUID(credentials.Username)); err == nil {
				if !crypt.HashAndEquals(credentials.Password, account.Password) {
					err = errors.New("password mismatch")
				} else {
					ret = account
				}
			}
			return
		})
}
