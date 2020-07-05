package parser

import (
	engine2 "go_crawler/crawler_base/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.7799520.com/jiaou/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine2.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine2.ParseResult{}
	limit := 10
	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine2.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})

		limit--
		if limit ==0 {
			 break
		}
	}
	return result
}
