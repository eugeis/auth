package auth

import (
	"context"
	"github.com/go-ee/utils/eh"
	"github.com/go-ee/utils/net"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/commandhandler/bus"
	"net/http"
)

type AccountHttpQueryHandler struct {
	*eh.HttpQueryHandler
	QueryRepository *AccountQueryRepository
}

func NewAccountHttpQueryHandlerFull(httpQueryHandler *eh.HttpQueryHandler, queryRepository *AccountQueryRepository) (ret *AccountHttpQueryHandler) {
	ret = &AccountHttpQueryHandler{
		HttpQueryHandler: httpQueryHandler,
		QueryRepository:  queryRepository,
	}
	return
}

func (o *AccountHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	ret, err := o.QueryRepository.FindAll()
	o.HandleResult(ret, err, "FindAllAccount", w, r)
}

func (o *AccountHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.Parse(vars["id"])
	ret, err := o.QueryRepository.FindById(id)
	o.HandleResult(ret, err, "FindByAccountId", w, r)
}

func (o *AccountHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request) {
	ret, err := o.QueryRepository.CountAll()
	o.HandleResult(ret, err, "CountAllAccount", w, r)
}

func (o *AccountHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.Parse(vars["id"])
	ret, err := o.QueryRepository.CountById(id)
	o.HandleResult(ret, err, "CountByAccountId", w, r)
}

func (o *AccountHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request) {
	ret, err := o.QueryRepository.ExistAll()
	o.HandleResult(ret, err, "ExistAllAccount", w, r)
}

func (o *AccountHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.Parse(vars["id"])
	ret, err := o.QueryRepository.ExistById(id)
	o.HandleResult(ret, err, "ExistByAccountId", w, r)
}

type AccountHttpCommandHandler struct {
	*eh.HttpCommandHandler
}

func NewAccountHttpCommandHandlerFull(httpCommandHandler *eh.HttpCommandHandler) (ret *AccountHttpCommandHandler) {
	ret = &AccountHttpCommandHandler{
		HttpCommandHandler: httpCommandHandler,
	}
	return
}

func (o *AccountHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.Parse(vars["id"])
	o.HandleCommand(&CreateAccount{Id: id}, w, r)
}

func (o *AccountHttpCommandHandler) Enable(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.Parse(vars["id"])
	o.HandleCommand(&EnableAccount{Id: id}, w, r)
}

func (o *AccountHttpCommandHandler) Disable(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.Parse(vars["id"])
	o.HandleCommand(&DisableAccount{Id: id}, w, r)
}

func (o *AccountHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.Parse(vars["id"])
	o.HandleCommand(&UpdateAccount{Id: id}, w, r)
}

func (o *AccountHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.Parse(vars["id"])
	o.HandleCommand(&DeleteAccount{Id: id}, w, r)
}

func (o *AccountHttpCommandHandler) SendEnabledConfirmation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.Parse(vars["id"])
	o.HandleCommand(&SendAccountEnabledConfirmation{Id: id}, w, r)
}

func (o *AccountHttpCommandHandler) SendDisabledConfirmation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.Parse(vars["id"])
	o.HandleCommand(&SendAccountDisabledConfirmation{Id: id}, w, r)
}

func (o *AccountHttpCommandHandler) Login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.Parse(vars["id"])
	o.HandleCommand(&LoginAccount{Id: id}, w, r)
}

func (o *AccountHttpCommandHandler) SendCreatedConfirmation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.Parse(vars["id"])
	o.HandleCommand(&SendAccountCreatedConfirmation{Id: id}, w, r)
}

type AccountRouter struct {
	PathPrefix     string
	QueryHandler   *AccountHttpQueryHandler
	CommandHandler *AccountHttpCommandHandler
	Router         *mux.Router
}

func NewAccountRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandHandler,
	readRepos func(string, func() (ret eventhorizon.Entity)) (ret eventhorizon.ReadWriteRepo)) (ret *AccountRouter) {
	pathPrefix = pathPrefix + "/" + "accounts"
	entityFactory := func() eventhorizon.Entity { return NewAccountDefault() }
	repo := readRepos(string(AccountAggregateType), entityFactory)
	httpQueryHandler := eh.NewHttpQueryHandlerFull()
	httpCommandHandler := eh.NewHttpCommandHandlerFull(context, commandBus)

	queryRepository := NewAccountQueryRepositoryFull(repo, context)
	queryHandler := NewAccountHttpQueryHandlerFull(httpQueryHandler, queryRepository)
	commandHandler := NewAccountHttpCommandHandlerFull(httpCommandHandler)
	ret = &AccountRouter{
		PathPrefix:     pathPrefix,
		QueryHandler:   queryHandler,
		CommandHandler: commandHandler,
	}
	return
}

func (o *AccountRouter) Setup(router *mux.Router) (err error) {
	router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
		Name("CountAccountById").HandlerFunc(o.QueryHandler.CountById).
		Queries(net.QueryType, net.QueryTypeCount)
	router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
		Name("CountAccountAll").HandlerFunc(o.QueryHandler.CountAll).
		Queries(net.QueryType, net.QueryTypeCount)
	router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
		Name("ExistAccountById").HandlerFunc(o.QueryHandler.ExistById).
		Queries(net.QueryType, net.QueryTypeExist)
	router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
		Name("ExistAccountAll").HandlerFunc(o.QueryHandler.ExistAll).
		Queries(net.QueryType, net.QueryTypeExist)
	router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
		Name("FindAccountById").HandlerFunc(o.QueryHandler.FindById)
	router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
		Name("FindAccountAll").HandlerFunc(o.QueryHandler.FindAll)
	router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
		Queries(net.Command, "login").
		Name("LoginAccount").HandlerFunc(o.CommandHandler.Login)
	router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
		Queries(net.Command, "sendCreatedConfirmation").
		Name("SendAccountCreatedConfirmation").HandlerFunc(o.CommandHandler.SendCreatedConfirmation)
	router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
		Queries(net.Command, "sendEnabledConfirmation").
		Name("SendAccountEnabledConfirmation").HandlerFunc(o.CommandHandler.SendEnabledConfirmation)
	router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
		Queries(net.Command, "sendDisabledConfirmation").
		Name("SendAccountDisabledConfirmation").HandlerFunc(o.CommandHandler.SendDisabledConfirmation)
	router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
		Name("CreateAccount").HandlerFunc(o.CommandHandler.Create)
	router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
		Queries(net.Command, "enable").
		Name("EnableAccount").HandlerFunc(o.CommandHandler.Enable)
	router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
		Queries(net.Command, "disable").
		Name("DisableAccount").HandlerFunc(o.CommandHandler.Disable)
	router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
		Name("UpdateAccount").HandlerFunc(o.CommandHandler.Update)
	router.Methods(http.MethodDelete).PathPrefix(o.PathPrefix).Path("/{id}").
		Name("DeleteAccount").HandlerFunc(o.CommandHandler.Delete)
	return
}

type AuthRouter struct {
	PathPrefix    string
	AccountRouter *AccountRouter
	Router        *mux.Router
}

func NewAuthRouter(pathPrefix string, context context.Context, commandBus *bus.CommandHandler,
	readRepos func(string, func() (ret eventhorizon.Entity)) (ret eventhorizon.ReadWriteRepo)) (ret *AuthRouter) {
	pathPrefix = pathPrefix + "/" + "auth"
	accountRouter := NewAccountRouter(pathPrefix, context, commandBus, readRepos)
	ret = &AuthRouter{
		PathPrefix:    pathPrefix,
		AccountRouter: accountRouter,
	}
	return
}

func (o *AuthRouter) Setup(router *mux.Router) (err error) {
	if err = o.AccountRouter.Setup(router); err != nil {
		return
	}
	return
}
