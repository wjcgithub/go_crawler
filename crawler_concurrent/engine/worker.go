package engine

import (
	fetcher2 "go_crawler/crawler_concurrent/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error)  {
	//log.Printf("Fetching %s", r.Url)
	body, err := fetcher2.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error "+
			"fetching url %s: %v",
			r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body, r.Url), nil
}
