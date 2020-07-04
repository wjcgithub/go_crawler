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

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}
