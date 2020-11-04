package models

type Response struct {
    Status int8 `json:"status"`
    Msg string `json:"msg"`
    Data []map[string]interface{} `json:"data"`
}

type Request struct {
    Values []map[string]interface{}
}


