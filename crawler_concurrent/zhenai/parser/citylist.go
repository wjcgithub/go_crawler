package parser

import (
	engine2 "go_crawler/crawler_concurrent/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.7799520.com/jiaou/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine2.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine2.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine2.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
