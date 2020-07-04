package parser

import (
	"imooc.com/ccmouse/learngo/crawler_concurrent/engine"
	"regexp"
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	const cityRe = `<a class="name" href="(http://www.7799520.com/user/[0-9]+\.html)"[^>]*>([^<]+)</a>`
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ProfileParser(string(m[2])),
		})
	}
	return result
}
