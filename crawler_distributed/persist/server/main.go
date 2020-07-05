package main

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	config2 "go_crawler/crawler_distributed/config"
	persist2 "go_crawler/crawler_distributed/persist"
	rpcsupport2 "go_crawler/crawler_distributed/rpcsupport"
	"log"
)

func main() {
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config2.ItemSaverPort) , config2.ElasticIndex))
}

func serveRpc(host, index string) error {
	client , err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return err
	}
	err = rpcsupport2.ServeRpc(host, &persist2.ItemSaverService{
		Client: client,
		Index:  index,
	})

	return err
}
