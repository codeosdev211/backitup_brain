package routes

import (
    "fmt"
    http "net/http"
    "encoding/json"
    models "../models"
    db "../db"
    fs "../filesys"
)

func AddFiles(res http.ResponseWriter, req *http.Request) {
    /* checking for request type */
    if req.Method != "POST" {
        json.NewEncoder(res).Encode(models.Response{1, "Invalid Request Type", nil})
    }
    var request models.FileRequest
    var response models.Response
    response.Status = 0
    response.Msg = "none"
    response.Data = nil

    /* validating request json object */
    err := json.NewDecoder(req.Body).Decode(&request)
    if err != nil {
        response.Status = 1
        response.Msg = "Invalid request body"
    }
    /* getting lists for fileInfo and fileData */
    fileList := request.Values[0].FileInfos
    dataList := request.Values[0].FileData

    var query string

    for iter, file := range fileList {
        filePath := fs.CreatePath(&file.OwnerCode, &file.Name)
        err = fs.WriteFile(&filePath, &dataList[iter])
        if err != nil {
            response.Status = 1
            response.Msg = "Could not write"
            break
        }

        query = "Update BAD set lastFileCode = lastFileCode + 1;"
        _, err = db.CallDatabase(false, &query)
        if err != nil {
            response.Status = 1
            response.Msg = "Could not update fileCode"
        }

        query = "select lastFileCode from BAD;"
        data, _ := db.CallDatabase(true, &query)
        fileCode := fmt.Sprintf("BUI%v", data[0]["lastFileCode"])
        query = fmt.Sprintf("Insert into BF(code, name, extension, originalSize, ownerCode, savedTo) values" +
            "('%v', '%v', '%v', '%v', '%v', '%v');",
            fileCode, file.Name, file.Extension, file.OriginalSize, file.OwnerCode, filePath)
        _, err = db.CallDatabase(false, &query)
        if err != nil {
            response.Status = 1
            response.Msg = "Database error"
        }
    }

    json.NewEncoder(res).Encode(response)
}


