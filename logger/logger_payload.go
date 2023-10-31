package logger

import "net/http"

type LoggerPayload struct {
	Name    string  `json:"name"`
	Payload Payload `json:"payload"`
}

type Payload struct {
	RequestId  string      `json:"requestId"`
	Time       string      `json:"time"`
	Header     http.Header `json:"header"`
	Url        string      `json:"url"`
	Endpoint   string      `json:"endpoint"`
	HttpMethod string      `json:"httpMethod"`
	Proto      string      `json:"proto"`
	UrlQuery   string      `json:"urlQuery"`
	Body       interface{} `json:"body"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
}
