package main

import (
	"imooc.com/ccmouse/learngo/crawler_concurrent/engine"
	"imooc.com/ccmouse/learngo/crawler_concurrent/persist"
	"imooc.com/ccmouse/learngo/crawler_concurrent/scheduler"
	"imooc.com/ccmouse/learngo/crawler_concurrent/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("wzly")
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
