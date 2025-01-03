package api

import (
	"github.com/imroc/req/v3"
	"github.com/solarhell/go-federalregister/timelocation"
)

type Client struct {
	httpClient *req.Client
}

func NewClient() *Client {
	timelocation.Setup()

	httpClient := req.NewClient()
	httpClient.SetCookieJar(nil)

	httpClient.EnableDumpAll()

	c := &Client{
		httpClient: httpClient,
	}

	return c
}
