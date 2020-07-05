package main

import (
	engine2 "go_crawler/crawler_concurrent/engine"
	persist2 "go_crawler/crawler_concurrent/persist"
	scheduler2 "go_crawler/crawler_concurrent/scheduler"
	parser2 "go_crawler/crawler_concurrent/zhenai/parser"
)

func main() {
	itemChan, err := persist2.ItemSaver("wzly")
	if err != nil {
		panic(err)
	}

	e := engine2.ConcurrentEngine{
		Scheduler:   &scheduler2.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	e.Run(engine2.Request{
		Url:        "http://www.7799520.com/jiaou",
		ParserFunc: parser2.ParseCityList,
	})
}
