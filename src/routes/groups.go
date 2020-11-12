package routes

import (
    "fmt"
    "time"
    http "net/http"
    _ "github.com/go-sql-driver/mysql"
    "encoding/json"
    models "../models"
    db "../db"
    helper "../commons"
)


func GetGroups(res http.ResponseWriter, req *http.Request) {
    /* checking for request type */
    if req.Method != "POST" {
        helper.SendErrorResponse(&res, "Invalid Request Type")
    }

    var request models.Request
    var response models.Response
    response.Status = 0
    response.Msg = "none"
    response.Data = nil

    /* validation request json object */
    err := json.NewDecoder(req.Body).Decode(&request)
    if err != nil {
        helper.SendErrorResponse(&res, "Invalid Request Body")
    }

    /* getting request body */
    //body := request.Values[0]

    // pending.....
}

func CreateGroup(res http.ResponseWriter, req *http.Request) {
    /* checking for request type */
    if req.Method != "POST" {
        helper.SendErrorResponse(&res, "Invalid Request Type")
    }

    var request models.Request
    var response models.Response
    response.Status = 0
    response.Msg = "none"
    response.Data = nil

    /* validating request json object */
    err := json.NewDecoder(req.Body).Decode(&request)
    if err != nil {
        helper.SendErrorResponse(&res, "Invalid Request Body")
    }

    /* getting request body */
    body := request.Values[0]
    /* updating list groupcode */
    query := "Update BAD set lastGroupCode = lastGroupCode + 1;"
    data, err := db.CallDatabase(false, &query)
    if err != nil {
        helper.SendErrorResponse(&res, "Database error")
    }

    /* creating new group */
    query = "select lastGroupCode from BAD;"
    data, err = db.CallDatabase(true, &query)
    if err != nil {
        helper.SendErrorResponse(&res, "Database error")
    }
    groupCode := fmt.Sprintf("BUI%v", data[0]["lastGroupCode"])
    currentTime := time.Now()
    query = fmt.Sprintf("insert into BG values(code, name, ownerCode, fileCount, userCount, createdOn) values " +
            "('%v', '%v', '%v', '%v', '%v', '%v');",
            groupCode, body["name"], body["ownerCode"], 0, 0, currentTime);
    _, err = db.CallDatabase(false, &query)
    if err != nil {
        helper.SendErrorResponse(&res, "Could not create group")
    }

    json.NewEncoder(res).Encode(response)
}



