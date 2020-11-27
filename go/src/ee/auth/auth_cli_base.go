package auth

import (
	"net/http"
)

type AccountCli struct {
	Client *AccountClient
}

func NewAccountCli(client *AccountClient) (ret *AccountCli) {
	ret = &AccountCli{
		Client: client,
	}
	return
}

type Cli struct {
	Client     *Client
	AccountCli *AccountCli
}

func NewCli(url string, httpClient *http.Client) (ret *Cli) {
	client := NewClient(url, httpClient)
	accountCli := NewAccountCli(client.AccountClient)
	ret = &Cli{
		Client:     client,
		AccountCli: accountCli,
	}
	return
}
