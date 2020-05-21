package engine

type Requests struct {
	Url       string
	ParseFunc func([]byte) ParseResult
	//结构体里加 函数
}

type ParseResult struct {
	Requests []Requests
	Items    []interface{}
	//结构体里加  接口
}
