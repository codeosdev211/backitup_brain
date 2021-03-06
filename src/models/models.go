package models

type Response struct {
    Status int8 `json:"status"`
    Msg string `json:"msg"`
    Data []map[string]interface{} `json:"data"`
}

type Request struct {
    Values []map[string]interface{} `json:"values"`
}

type BF struct {
    Id int `json:"id"`
    Code string `json:"code"`
    Name string `json:"name"`
    Extension string `json:"extension"`
    OriginalSize string `json:"originalSize"`
    OwnerCode string `json:"ownerCode"`
    SavedTo string `json:"savedTo"`
}

type FileInfo struct {
    FileInfos []BF
    FileData []string
}

type FileRequest struct {
    Values []FileInfo
}

type FileResponse struct {
    Status int8 `json:"status"`
    Msg string `json:"msg"`
    FileName string `json:"fileName"`
    FileData string `json:"fileData"`
}

type GroupResponse struct {
    Status int8 `json:"status"`
    Msg string `json:"msg"`
    UserGroups []map[string]interface{} `json:"userGroups"`
    PublicGroups []map[string]interface{} `json:"publicGroups"`
}


