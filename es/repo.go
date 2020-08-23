package es

type Repo struct {
	client      *Client
	mappingFile string
}

func (Repo) Foo() {}

func NewRepo(client *Client, mappingFile string) *Repo {
	return &Repo{
		client:      client,
		mappingFile: mappingFile,
	}
}
