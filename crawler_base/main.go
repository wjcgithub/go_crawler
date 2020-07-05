package main

import (
	engine2 "go_crawler/crawler_base/engine"
	parser2 "go_crawler/crawler_base/zhenai/parser"
)

func main() {
	engine2.Run(engine2.Request{
		Url:        "http://www.7799520.com/jiaou",
		ParserFunc: parser2.ParseCityList,
	})
}
