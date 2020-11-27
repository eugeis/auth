package auth

import (
	"encoding/json"
	"github.com/go-ee/utils/net"
	"io/ioutil"
	"net/http"
)

type AccountClient struct {
	Url    string
	Client *http.Client
}

func NewAccountClient(url string, client *http.Client) (ret *AccountClient) {
	url = url + "/" + "accounts"
	ret = &AccountClient{
		Url:    url,
		Client: client,
	}
	return
}

func (o *AccountClient) ImportJSON(fileJSON string) (err error) {
	var items []*Account
	if items, err = o.ReadFileJSON(fileJSON); err != nil {
		return
	}

	err = o.Create(items)
	return
}

func (o *AccountClient) Create(items []*Account) (err error) {
	for _, item := range items {
		if err = net.PostById(item, item.Id, o.Url, o.Client); err != nil {
			return
		}
	}
	return
}

func (o *AccountClient) ReadFileJSON(fileJSON string) (ret []*Account, err error) {
	jsonBytes, _ := ioutil.ReadFile(fileJSON)

	err = json.Unmarshal(jsonBytes, &ret)
	return
}

type Client struct {
	Url           string
	Client        *http.Client
	AccountClient *AccountClient
}

func NewClient(url string, client *http.Client) (ret *Client) {
	url = url + "/" + "auth"
	accountClient := NewAccountClient(url, client)
	ret = &Client{
		Url:           url,
		Client:        client,
		AccountClient: accountClient,
	}
	return
}
