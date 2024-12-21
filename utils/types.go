package utils

type Response struct {
	Error string `json:"error"`
}

type Answer struct {
	Result string `json:"result"`
}

type Expression struct {
	Expression string 
}