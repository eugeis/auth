package auth

import (
	"github.com/google/uuid"
	"github.com/urfave/cli"
	"net/http"
	"strings"
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

func (o *AccountCli) BuildCommands() (ret []cli.Command) {
	ret = []cli.Command{
		o.BuildCommandImportJSON(), o.BuildCommandExportJSON(), o.BuildCommandDeleteById(), o.BuildCommandDeleteByIds(),
	}

	return
}

func (o *AccountCli) BuildCommandImportJSON() (ret cli.Command) {

	return
}

func (o *AccountCli) BuildCommandExportJSON() (ret cli.Command) {

	return
}

func (o *AccountCli) BuildCommandDeleteByIds() (ret cli.Command) {
	ret = cli.Command{
		Name:  "deleteByIds",
		Usage: "delete Account by ids",
		Flags: []cli.Flag{&cli.StringFlag{
			Name:     "ids",
			Usage:    "ids of the Accounts to delete, separated by semicolon",
			Required: true,
		}},
		Action: func(c *cli.Context) (err error) {
			var id uuid.UUID
			var ids []uuid.UUID
			for _, idString := range strings.Split(c.String("ids"), ",") {
				if id, err = uuid.Parse(idString); err != nil {
					return
				}
				ids = append(ids, id)
			}
			err = o.Client.DeleteByIds(ids)
			return
		},
	}
	return
}

func (o *AccountCli) BuildCommandDeleteById() (ret cli.Command) {
	ret = cli.Command{
		Name:  "deleteById",
		Usage: "delete Account by id",
		Flags: []cli.Flag{&cli.StringFlag{
			Name:     "id",
			Usage:    "id of the Account to delete",
			Required: true,
		}},
		Action: func(c *cli.Context) (err error) {
			var id uuid.UUID
			if id, err = uuid.Parse(c.String("id")); err == nil {
				err = o.Client.DeleteById(&id)
			}
			return
		},
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
