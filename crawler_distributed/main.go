package main

import (
	"fmt"
	"go_crawler/crawler_concurrent/engine"
	"go_crawler/crawler_concurrent/scheduler"
	"go_crawler/crawler_concurrent/zhenai/parser"
	config2 "go_crawler/crawler_distributed/config"
	client2 "go_crawler/crawler_distributed/persist/client"
)

func main() {
	itemChan, err := client2.ItemSaver(fmt.Sprintf(":%d", config2.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		Url:        "http://www.7799520.com/jiaou",
		ParserFunc: parser.ParseCityList,
	})
}
