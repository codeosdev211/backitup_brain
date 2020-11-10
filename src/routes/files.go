package routes

import (
    "fmt"
    http "net/http"
    "encoding/json"
    models "../models"
    _ "github.com/go-sql-driver/mysql"
    db "../db"
    fs "../filesys"
    helper "../commons"
)

func AddFiles(res http.ResponseWriter, req *http.Request) {
    /* checking for request type */
    if req.Method != "POST" {
        helper.SendErrorResponse(&res, "Invalid Request Type")
    }
    var request models.FileRequest
    var response models.Response
    response.Status = 0
    response.Msg = "none"
    response.Data = nil

    /* validating request json object */
    err := json.NewDecoder(req.Body).Decode(&request)
    if err != nil {
        helper.SendErrorResponse(&res, "Invalid Request Body")
    }
    /* getting lists for fileInfo and fileData */
    fileList := request.Values[0].FileInfos
    dataList := request.Values[0].FileData

    var query string

    for iter, file := range fileList {
        filePath := fs.CreatePath(&file.OwnerCode, &file.Name)
        err = fs.WriteFile(&filePath, &dataList[iter])
        if err != nil {
            helper.SendErrorResponse(&res, "Could not write file(s)")
        }

        query = "Update BAD set lastFileCode = lastFileCode + 1;"
        _, err = db.CallDatabase(false, &query)
        if err != nil {
            helper.SendErrorResponse(&res, "Could not update file Code")
        }

        query = "select lastFileCode from BAD;"
        data, _ := db.CallDatabase(true, &query)
        fileCode := fmt.Sprintf("BUI%v", data[0]["lastFileCode"])
        query = fmt.Sprintf("Insert into BF(code, name, extension, originalSize, ownerCode, savedTo) values" +
            "('%v', '%v', '%v', '%v', '%v', '%v');",
            fileCode, file.Name, file.Extension, file.OriginalSize, file.OwnerCode, filePath)
        _, err = db.CallDatabase(false, &query)
        if err != nil {
            helper.SendErrorResponse(&res, "Database error")
        }
        query = fmt.Sprintf("update BU set totalFiles = totalFiles + 1 where code='%v';", file.OwnerCode)
        _, err = db.CallDatabase(false, &query)
        if err != nil {
            helper.SendErrorResponse(&res, "Database error")
        }
    }

    json.NewEncoder(res).Encode(response)
}

func ListFiles(res http.ResponseWriter, req *http.Request) {
    /* checking for request type */
    if req.Method != "POST" {
        helper.SendErrorResponse(&res, "Invalid request type")
    }
    var request models.Request
    var response models.Response
    response.Status = 0
    response.Msg = "none"
    response.Data = nil

    /* validating request json object */
    err := json.NewDecoder(req.Body).Decode(&request)
    if err != nil {
        helper.SendErrorResponse(&res, "Invalid request Body")
    }

    body := request.Values[0]
    query := fmt.Sprintf("select * from BF where ownerCode='%v';", body["code"])
    response.Data, err = db.CallDatabase(true, &query)
    if err != nil {
        helper.SendErrorResponse(&res, "Database error")
    }

    json.NewEncoder(res).Encode(response)
}

func DownloadFile(res http.ResponseWriter, req *http.Request) {
    /* checking for request type */
    if req.Method != "POST" {
        helper.SendErrorResponse(&res, "Invalid request type")
    }
    var request models.Request
    var response models.FileResponse
    response.Status = 0
    response.Msg = "none"

    /* validating request json object */
    err := json.NewDecoder(req.Body).Decode(&request)
    if err != nil {
        helper.SendErrorResponse(&res, "Invalid request body")
    }

    body := request.Values[0]
    query :=  fmt.Sprintf("select name, savedTo from BF where code='%v' and ownerCode='%v';",
            body["code"], body["ownerCode"])
    data, err := db.CallDatabase(true, &query)
    if err != nil {
        helper.SendErrorResponse(&res, "Database error")
    }
    path := fmt.Sprintf("%v", data[0]["savedTo"])
    file, err := fs.ReadFile(&path)
    if err != nil {
        helper.SendErrorResponse(&res, "Could not read file")
    }

    response.FileName = fmt.Sprintf("%v", data[0]["name"])
    response.FileData = fmt.Sprintf("%v", file)

    json.NewEncoder(res).Encode(response)
}
