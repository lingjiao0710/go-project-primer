package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

//空的parser，什么都不做
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
