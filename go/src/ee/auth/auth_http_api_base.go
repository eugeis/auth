package auth

import (
	"context"
	"github.com/go-ee/utils/eh"
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
	o.HandleResult(ret, err, "AccountFindAll", w, r)
}

func (o *AccountHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.Parse(vars["id"])
	ret, err := o.QueryRepository.FindById(id)
	o.HandleResult(ret, err, "AccountFindById", w, r)
}

func (o *AccountHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request) {
	ret, err := o.QueryRepository.CountAll()
	o.HandleResult(ret, err, "AccountCountAll", w, r)
}

func (o *AccountHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.Parse(vars["id"])
	ret, err := o.QueryRepository.CountById(id)
	o.HandleResult(ret, err, "AccountCountById", w, r)
}

func (o *AccountHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request) {
	ret, err := o.QueryRepository.ExistAll()
	o.HandleResult(ret, err, "AccountExistAll", w, r)
}

func (o *AccountHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.Parse(vars["id"])
	ret, err := o.QueryRepository.ExistById(id)
	o.HandleResult(ret, err, "AccountExistById", w, r)
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
	o.HandleCommand(&SendEnabledConfirmationAccount{Id: id}, w, r)
}

func (o *AccountHttpCommandHandler) SendDisabledConfirmation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.Parse(vars["id"])
	o.HandleCommand(&SendDisabledConfirmationAccount{Id: id}, w, r)
}

func (o *AccountHttpCommandHandler) Login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.Parse(vars["id"])
	o.HandleCommand(&LoginAccount{Id: id}, w, r)
}

type AccountRouter struct {
	PathPrefix        string
	PathPrefixIdBased string
	QueryHandler      *AccountHttpQueryHandler
	CommandHandler    *AccountHttpCommandHandler
}

func NewAccountRouter(pathPrefix string, newContext func(string) (ret context.Context), commandBus *bus.CommandHandler,
	repo eventhorizon.ReadRepo) (ret *AccountRouter) {
	pathPrefixIdBased := pathPrefix + "/" + "account"
	pathPrefix = pathPrefix + "/" + "accounts"
	ctx := newContext("account")
	httpQueryHandler := eh.NewHttpQueryHandlerFull()
	httpCommandHandler := eh.NewHttpCommandHandlerFull(ctx, commandBus)

	queryRepository := NewAccountQueryRepositoryFull(repo, ctx)
	queryHandler := NewAccountHttpQueryHandlerFull(httpQueryHandler, queryRepository)
	commandHandler := NewAccountHttpCommandHandlerFull(httpCommandHandler)
	ret = &AccountRouter{
		PathPrefix:        pathPrefix,
		PathPrefixIdBased: pathPrefixIdBased,
		QueryHandler:      queryHandler,
		CommandHandler:    commandHandler,
	}
	return
}

func (o *AccountRouter) Setup(router *mux.Router) (err error) {
	router.Methods(http.MethodGet).PathPrefix(o.PathPrefixIdBased).Path("/{id}").
		Name("AccountFindById").
		HandlerFunc(o.QueryHandler.FindById)
	router.Methods(http.MethodGet).PathPrefix(o.PathPrefixIdBased).Path("/{id}/count").
		Name("AccountCountById").
		HandlerFunc(o.QueryHandler.CountById)
	router.Methods(http.MethodGet).PathPrefix(o.PathPrefixIdBased).Path("/{id}/exist").
		Name("AccountExistById").
		HandlerFunc(o.QueryHandler.ExistById)
	router.Methods(http.MethodPost).PathPrefix(o.PathPrefixIdBased).Path("/{id}").
		Name("CreateAccount").
		HandlerFunc(o.CommandHandler.Create)
	router.Methods(http.MethodPost).PathPrefix(o.PathPrefixIdBased).Path("/{id}/login").
		Name("LoginAccount").
		HandlerFunc(o.CommandHandler.Login)
	router.Methods(http.MethodPost).PathPrefix(o.PathPrefixIdBased).Path("/{id}/send-enabled-confirmation").
		Name("SendEnabledConfirmationAccount").
		HandlerFunc(o.CommandHandler.SendEnabledConfirmation)
	router.Methods(http.MethodPost).PathPrefix(o.PathPrefixIdBased).Path("/{id}/send-disabled-confirmation").
		Name("SendDisabledConfirmationAccount").
		HandlerFunc(o.CommandHandler.SendDisabledConfirmation)
	router.Methods(http.MethodPut).PathPrefix(o.PathPrefixIdBased).Path("/{id}").
		Name("UpdateAccount").
		HandlerFunc(o.CommandHandler.Update)
	router.Methods(http.MethodPut).PathPrefix(o.PathPrefixIdBased).Path("/{id}/enable").
		Name("EnableAccount").
		HandlerFunc(o.CommandHandler.Enable)
	router.Methods(http.MethodPut).PathPrefix(o.PathPrefixIdBased).Path("/{id}/disable").
		Name("DisableAccount").
		HandlerFunc(o.CommandHandler.Disable)
	router.Methods(http.MethodDelete).PathPrefix(o.PathPrefixIdBased).Path("/{id}").
		Name("DeleteAccount").
		HandlerFunc(o.CommandHandler.Delete)
	router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("").
		Name("AccountFindAll").
		HandlerFunc(o.QueryHandler.FindAll)
	router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/count").
		Name("AccountCountAll").
		HandlerFunc(o.QueryHandler.CountAll)
	router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/exist").
		Name("AccountExistAll").
		HandlerFunc(o.QueryHandler.ExistAll)
	return
}

type Router struct {
	PathPrefix    string
	AccountRouter *AccountRouter
}

func NewRouter(pathPrefix string, newContext func(string) (ret context.Context), esEngine *EsEngine) (ret *Router, err error) {
	pathPrefix = pathPrefix + "/" + "auth"

	var projectorAccount *AccountProjector
	if projectorAccount, err = esEngine.Account.RegisterAccountProjector(string(AccountAggregateType),
		esEngine.Account.AggregateHandlers, esEngine.Account.Events); err != nil {
		return
	}

	accountRouter := NewAccountRouter(pathPrefix, newContext, esEngine.CommandBus, projectorAccount.Repo)

	ret = &Router{
		PathPrefix:    pathPrefix,
		AccountRouter: accountRouter,
	}
	return
}

func (o *Router) Setup(router *mux.Router) (err error) {
	if err = o.AccountRouter.Setup(router); err != nil {
		return
	}
	return
}
