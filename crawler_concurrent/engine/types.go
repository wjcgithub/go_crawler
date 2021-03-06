package engine

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type ParserFunc func(contents []byte, url string) ParseResult

type Request struct {
	Url        string
	ParserFunc ParserFunc
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}
