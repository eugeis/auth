package auth

import (
	"encoding/json"
	"github.com/go-ee/utils/net"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
)

type AccountClient struct {
	UrlIdBased string
	Url        string
	Client     *http.Client
}

func NewAccountClient(url string, client *http.Client) (ret *AccountClient) {
	urlIdBased := url + "/" + "account"
	url = url + "/" + "accounts"
	ret = &AccountClient{
		UrlIdBased: urlIdBased,
		Url:        url,
		Client:     client,
	}
	return
}

func (o *AccountClient) ImportJSON(fileJSON string) (err error) {
	var items []*Account
	if items, err = o.ReadFileJSON(fileJSON); err != nil {
		return
	}

	err = o.CreateItems(items)
	return
}

func (o *AccountClient) CreateItems(items []*Account) (err error) {
	for _, item := range items {
		if err = net.PostById(item, item.Id, o.UrlIdBased, o.Client); err != nil {
			return
		}
	}
	return
}

func (o *AccountClient) DeleteItems(items []*Account) (err error) {
	for _, item := range items {
		if err = net.DeleteById(item.Id, o.UrlIdBased, o.Client); err != nil {
			return
		}
	}
	return
}

func (o *AccountClient) DeleteById(itemId *uuid.UUID) (err error) {
	err = net.DeleteById(itemId, o.UrlIdBased, o.Client)
	return
}

func (o *AccountClient) FindAll() (ret []*Account, err error) {
	err = net.GetItems(&ret, o.Url, o.Client)
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
