package main

import (
	"fmt"

	"demo-viper/conf"
	"demo-viper/curlec"
	"demo-viper/es"
)

func main() {
	conf.Init()
	fmt.Print(conf.R)

	curlecClient := curlec.NewClient(conf.R.Curlec)
	curlecClient.Foo()

	esClient := es.NewClient(conf.R.Elasticsearch.Host)
	mandateRepo := es.NewRepo(esClient, conf.R.Elasticsearch.Repository.Mapping["mandate"])
	mandateRepo.Foo()

	select {}

}
