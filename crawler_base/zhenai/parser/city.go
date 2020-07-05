package parser

import (
	engine2 "go_crawler/crawler_base/engine"
	"regexp"
)

func ParseCity(contents []byte) engine2.ParseResult {
	const cityRe = `<a class="name" href="(http://www.7799520.com/user/[0-9]+\.html)"[^>]*>([^<]+)</a>`
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine2.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine2.Request{
			Url:        string(m[1]),
			ParserFunc: func(c []byte) engine2.ParseResult {
				return ParseProfile(c, name)
			},
		})
	}
	return result
}
