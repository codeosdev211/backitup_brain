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


func GetStats(res http.ResponseWriter, req *http.Request) {
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
    query := fmt.Sprintf("select totalFiles, totalGroups from BU where code='%v';", body["code"])
    response.Data, err = db.CallDatabase(true, &query)
    if err != nil {
        helper.SendErrorResponse(&res, "Database error")
    }
    json.NewEncoder(res).Encode(response)
}

func SignInUser(res http.ResponseWriter, req *http.Request) {
    /* checking for request Type */
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
    /* checking for existing user */
    var query string = fmt.Sprintf("select count(*) as isThere from BU where email='%v' and password='%v';", body["email"], body["password"])
    data, err := db.CallDatabase(true, &query)
    if err != nil {
        helper.SendErrorResponse(&res, "Database error")
    }
    if data[0]["isThere"] == "0" {
        helper.SendErrorResponse(&res, "Invalid email or password")
    }else{
        query = fmt.Sprintf("select * from BU where email='%s' and password='%s';", body["email"], body["password"])
        response.Data, err = db.CallDatabase(true, &query)
        if err != nil {
            helper.SendErrorResponse(&res, "Could not get user data")
        }
    }

    json.NewEncoder(res).Encode(response)
}

func SignUpUser(res http.ResponseWriter, req *http.Request) {
    /* checking for request Type */
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
        helper.SendErrorResponse(&res, "Invalid request body")
    }


    body := request.Values[0]
    /* checking for existing user */
    var query string = fmt.Sprintf("select count(*) as isThere from BU where email='%v' and password='%v';", body["email"], body["password"])
    data, err := db.CallDatabase(true, &query)
    if err != nil {
        helper.SendErrorResponse(&res, "Database error")
    }

    if data[0]["isThere"] != "0" {
        helper.SendErrorResponse(&res, "User already exists")
    }else{
        /* updating last usercode*/
        query = "Update BAD set lastUserCode = lastUserCode + 1;"
        _, err = db.CallDatabase(false, &query)
        if err != nil {
            helper.SendErrorResponse(&res, "Database error")
        }

        /*creating new user */
        query = "select lastUserCode from BAD;"
        data, err = db.CallDatabase(true, &query)
        if err != nil {
            helper.SendErrorResponse(&res, "Database error")
        }
	    userCode := fmt.Sprintf("BUI%v", data[0]["lastUserCode"])
        currentTime := time.Now()
        query = fmt.Sprintf("Insert into BU (code, firstName, lastName, email, password, totalGroups, totalFiles, createdOn, isActive) values " +
            "('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v');",
            userCode, body["firstName"], body["lastName"], body["email"], body["password"], 0, 0, currentTime.Format("2006.01.02 15:04:05"), "TRUE");
        _, err = db.CallDatabase(false, &query)
        if err != nil {
            helper.SendErrorResponse(&res, "Could not create user")
        }
    }

    json.NewEncoder(res).Encode(response)
}

