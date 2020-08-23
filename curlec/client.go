package curlec

import (
	"demo-viper/conf"
)

type Client struct {
	cfg conf.CurlecConf
}

func (Client) Foo() {}

func NewClient(cfg conf.CurlecConf) *Client {
	return &Client{cfg: cfg}
}
