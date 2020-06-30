package main

import (
	"imooc.com/ccmouse/learngo/crawler_concurrent/engine"
	"imooc.com/ccmouse/learngo/crawler_concurrent/persist"
	"imooc.com/ccmouse/learngo/crawler_concurrent/scheduler"
	"imooc.com/ccmouse/learngo/crawler_concurrent/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        "http://www.7799520.com/jiaou",
		ParserFunc: parser.ParseCityList,
	})
}
