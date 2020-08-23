package curlec

import (
	"demo-viper/conf"
	"fmt"
)

type Client struct {
	cfg *conf.CurlecConf
}

func (c Client) Foo() {
	fmt.Println(c.cfg.Host)
}

func NewClient(cfg *conf.CurlecConf) *Client {
	return &Client{cfg: cfg}
}
