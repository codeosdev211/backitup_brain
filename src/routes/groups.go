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

func MembersList(res http.ResponseWriter, req *http.Request) {
    /* checking for request type */
    if req.Method != "POST" {
        helper.SendErrorResponse(&res, "Invalid Request Type")
    }

    var request models.Request
    var response models.Response
    response.Status = 0
    response.Msg = "none"
    response.Data =  nil

    /*validating request json object */
    err := json.NewDecoder(req.Body).Decode(&request)
    if err != nil {
        helper.SendErrorResponse(&res, "Invalid Request Body")
    }
    //body := request.Values[0]

}


func GetGroups(res http.ResponseWriter, req *http.Request) {
    /* checking for request type */
    if req.Method != "POST" {
        helper.SendErrorResponse(&res, "Invalid Request Type")
    }

    var request models.Request
    var response models.GroupResponse
    response.Status = 0
    response.Msg = "none"
    response.UserGroups = nil
    response.PublicGroups = nil

    /* validation request json object */
    err := json.NewDecoder(req.Body).Decode(&request)
    if err != nil {
        helper.SendErrorResponse(&res, "Invalid Request Body")
    }

    body := request.Values[0]

    query := fmt.Sprintf("select BG.code, BG.name, BG.ownerCode, BG.userCount, BU.firstName from BG left join BU on BG.code like '%%';")
    response.PublicGroups, err = db.CallDatabase(true, &query)
    if err != nil {
        helper.SendErrorResponse(&res, "Database error")
    }
    /*
     Now... i figured out a way to get the list,
     joining THREE tables should do, right? ... YES!
    */

    query = fmt.Sprintf("select BG.code, BG.name, BG.userCount, BU.firstName from BUG right join BG "+
                        " right join BU on BG.ownerCode = BU.code on BUG.groupCode = BG.code where BUG.userCode='%v' ;",
                        body["code"])
    response.UserGroups, err = db.CallDatabase(true, &query)
    if err != nil {
        helper.SendErrorResponse(&res, "Database error")
    }

    json.NewEncoder(res).Encode(response)
}

func FilesShared(res http.ResponseWriter, req *http.Request) {
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

    /* now for those who did not understand the query...
       its basically BFG has all of BF and BF has all of BU
       so its BFG -> BF -> BU
       and not BU <- BFG -> BF 
    */
    query :=  fmt.Sprintf("select BF.code, BF.name, BF.ownerCode, BU.firstName from BFG right join BF " +
                          "right join BU on BF.ownerCode = BU.code on BFG.fileCode = BF.code where BFG.groupCode='%v';",
                          body["code"])
    response.Data, err = db.CallDatabase(true, &query)
    if err != nil {
        helper.SendErrorResponse(&res, "Database error")
    }

    json.NewEncoder(res).Encode(response)
}

func ShareFile(res http.ResponseWriter, req *http.Request) {
    /* checking for request type */
    if req.Method != "POST" {
        helper.SendErrorResponse(&res, "Invalid request Type")
    }
    var request models.Request
    var response models.Response
    response.Status = 0
    response.Msg = "none"
    response.Data = nil

    /* validating request json object */
    err := json.NewDecoder(req.Body).Decode(&request)
    if err != nil {
        helper.SendErrorResponse(&res, "Invalid Request body")
    }

    /* getting request body */
    body := request.Values[0]

    currentTime := time.Now()
    query := fmt.Sprintf("insert into BFG(fileCode, groupCode, addedOn, addedBy) values " +
                        " ('%v', '%v', '%v', '%v');",
                        body["fileCode"], body["groupCode"], currentTime.Format("2006.01.02 15:04:05"), body["addedBy"])
    _, err = db.CallDatabase(false, &query)
    if err != nil {
        helper.SendErrorResponse(&res, "Database error")
    }

    json.NewEncoder(res).Encode(response)
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
    query = fmt.Sprintf("insert into BG(code, name, ownerCode, fileCount, userCount, createdOn) values " +
            "('%v', '%v', '%v', '%v', '%v', '%v');",
            groupCode, body["name"], body["ownerCode"], 0, 0, currentTime.Format("2006.01.02 15:04:05"))
    _, err = db.CallDatabase(false, &query)
    if err != nil {
        helper.SendErrorResponse(&res, "Could not create group")
    }

    query = "update BU set totalGroups = totalGroups + 1;"
    _, err = db.CallDatabase(false, &query)
    if err != nil {
        helper.SendErrorResponse(&res, "Could not update group count")
    }

    query = fmt.Sprintf("Insert into BUG(userCode, groupCode, addedOn, addedBy) values "+
            "('%v', '%v', '%v', '%v');",
            body["ownerCode"], groupCode, currentTime.Format("2006.01.02 15:04:05"), body["ownerCode"])
    _, err = db.CallDatabase(false, &query)
    if err != nil {
        helper.SendErrorResponse(&res, "Could not log group creation")
    }

    json.NewEncoder(res).Encode(response)
}



