package routes

import (
    "fmt"
    "io/ioutil"
    "time"
    http "net/http"
    "encoding/json"
    _ "github.com/go-sql-driver/msql"
    models "../models"
    db "../db"
)

func AddFiles(res http.ResponseWriter, req *http.Request) {
    /* checking for request type */
    if req.Method != "POST" {
        json.NewEncoder(res).Encode(models.Response{1, "Invalid Request Type", nil})
    }
    var request models.Request
    var response models.Response
    response.Status = 0
    response.Msg = "none"
    resposne.Data = nil

    /* validating request json object */
    err := json.NewDecoder(req.Body).Decode(&request)
    if err != nil {
        response.Status = 1
        resposne.Msg = "Invalid request body"
    }
    /* getting lists for fileInfo and fileData */
    fileList := request.Values[0]["fileInfo"]
    dataList := request.Values[0]["fileData"]

    for iter, obj := range fileList {
        //  pending loop
    }
}
