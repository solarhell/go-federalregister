package federalregister

import (
	"github.com/imroc/req/v3"
	"github.com/solarhell/go-federalregister/timelocation"
)

type Client struct {
	httpClient *req.Client
}

func NewClient(reqClient *req.Client) *Client {
	timelocation.Setup()

	httpClient := reqClient
	if httpClient == nil {
		httpClient = req.NewClient()
		httpClient.SetCookieJar(nil)
	}

	c := &Client{
		httpClient: httpClient,
	}

	return c
}
