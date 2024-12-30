package utils

import "io"

type Response struct {
	Error string `json:"error"`
}

type Answer struct {
	Result string `json:"result"`
}

type Expression struct {
	Expression string 
}

type TestRequest struct {
	Method 		string
	Name 		string
	Expression 	io.Reader
	StatusCode 	int
}