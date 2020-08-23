package es

type Client struct {
	host string
}

func NewClient(host string) *Client {
	return &Client{host: host}
}
