package routes

import (
    "fmt"
    "time"
    http "net/http"
    _ "github.com/go-sql-driver/mysql"
    "encoding/json"
    models "../models"
    db "../db"
)


func SignInUser(res http.ResponseWriter, req *http.Request) {
    /* checking for request Type */
    if req.Method != "POST" {
        json.NewEncoder(res).Encode(models.Response{1, "Invalid Request Type", nil})
    }
    var request models.Request
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
    /* getting request body */
    body := request.Values[0]
    /* checking for existing user */
    var query string = fmt.Sprintf("select count(*) as isThere from BU where email='%v' and password='%v';", body["email"], body["password"])
    data, err := db.CallDatabase(true, &query)
    if err != nil {
        response.Status = 1
        response.Msg = "Database error"
    }
    if data[0]["isThere"] == "0" {
        response.Status = 1
        response.Msg = "Invalid Email or Password"
    }else{
        query = fmt.Sprintf("select * from BU where email='%s' and password='%s';", body["email"], body["password"])
        response.Data, err = db.CallDatabase(true, &query)
        if err != nil {
            response.Status = 1
            response.Msg = "Could not get user data"
        }
    }

    json.NewEncoder(res).Encode(response)
}

func SignUpUser(res http.ResponseWriter, req *http.Request) {
    /* checking for request Type */
    if req.Method != "POST" {
        json.NewEncoder(res).Encode(models.Response{1, "Invalid Request Type", nil})
    }
    var request models.Request
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


    body := request.Values[0]
    /* checking for existing user */
    var query string = fmt.Sprintf("select count(*) as isThere from BU where email='%v' and password='%v';", body["email"], body["password"])
    data, err := db.CallDatabase(true, &query)
    if err != nil {
        response.Status = 1
        response.Msg = "Database error"
    }

    if data[0]["isThere"] != "0" {
        response.Status = 1
        response.Msg = "User already exists"
    }else{
        /* updating last usercode*/
        query = "Update BAD set lastUserCode = lastUserCode + 1;"
        _, err = db.CallDatabase(false, &query)
        if err != nil {
            response.Status = 1
            response.Msg = "Database error"
        }

        /*creating new user */
        query = "select lastUserCode from BAD;"
        data, err = db.CallDatabase(true, &query)
        if err != nil {
            panic(err)
        }
	    userCode := fmt.Sprintf("BUI%v", data[0]["lastUserCode"])
        currentTime := time.Now()
        query = fmt.Sprintf("Insert into BU (code, firstName, lastName, email, password, totalGroups, totalFiles, createdOn, isActive) values " +
            "('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v');",
            userCode, body["firstName"], body["lastName"], body["email"], body["password"], 0, 0, currentTime.Format("2006.01.02 15:04:05"), "TRUE");
        _, err = db.CallDatabase(false, &query)
        if err != nil {
            response.Status = 1
            response.Msg = "Could not create User"
        }
    }

    json.NewEncoder(res).Encode(response)
}
