package web

type Request struct {
	ID   string
	Body *RequestBody
	Err  error
}

type RequestBody struct {
	Params    RequestParams `json:"-"`
	RequestID string        `json:"-"`
}

type RequestParams struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}
