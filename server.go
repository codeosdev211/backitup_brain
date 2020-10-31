package main

import (
    "fmt"
    http "net/http"
    "encoding/json"
    mux "github.com/gorilla/mux"
    db "./db"
    _ "github.com/go-sql-driver/mysql"
    "time"
)

type Response struct {
    Status int8 `json:"status"`
    Msg string `json:"msg"`
    Data []map[string]interface{} `json:"data"`
}

type Request struct {
    Values []map[string]interface{}
}

/* main function allows routing requests */
func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/home", homePage)
    router.HandleFunc("/user/login", signInUser)
    router.HandleFunc("/user/register", signUpUser)

    http.ListenAndServe(":8001", router)
}

/* setting headers for CORS */
func setupCORS(response *http.ResponseWriter) {
    (*response).Header().Set("Access-Control-Allow-Origin", "*")
    (*response).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*response).Header().Set("Access-Control-Allow-Headers", "*")
}


/* homepage function */
func homePage(res http.ResponseWriter, req *http.Request) {
    var status int8 = 0
    var msg string = "none"
    setupCORS(&res)
    var response Response
    response.Status = status
    response.Msg = msg
    json.NewEncoder(res).Encode(response)
}

/* signin */
func signInUser(res http.ResponseWriter, req *http.Request) {
    if req.Method != "POST" {
        json.NewEncoder(res).Encode(Response{1, "Invalid Request Type", nil})
    }
    var request Request 
    var msg string = "none"
    var status int8 = 0

    err := json.NewDecoder(req.Body).Decode(&request)
    if err != nil {
        status = 1
        msg = "Invalid request body"
    }
    body := request.Values[0]
    var query string = fmt.Sprintf("select count(*) as isThere from BU where email='%v' and password='%v';", body["email"], body["password"])

    data, err := db.CallDatabase(true, &query)
    if err != nil {
        status = 1
        msg = "Database error"
    }

    if data[0]["isThere"] == 0 {
        status = 1
        msg = "Invalid Email or Password"
    }else{
        query = fmt.Sprintf("select * from BU where email='%s' and password='%s';", body["email"], body["password"])

        data , err = db.CallDatabase(true, &query)
        if err != nil {
            status = 1
            msg = "Could not get user data"
        }
    }
    var response Response
    response.Status = status
    response.Msg = msg
    response.Data = data

    json.NewEncoder(res).Encode(response)
}


func signUpUser(res http.ResponseWriter, req *http.Request) {
    if req.Method != "POST" {
        json.NewEncoder(res).Encode(Response{1, "Invalid Request Type", nil})
    }
    var request Request
    var msg string = "none"
    var status int8 = 0

    err := json.NewDecoder(req.Body).Decode(&request)
    if err != nil {
        status = 1
        msg = "Invalid request body"
    }
    body := request.Values[0]
    var query string = fmt.Sprintf("select count(*) as isThere from BU where email='%v' and password='%v';", body["email"], body["password"])
    data, err := db.CallDatabase(true, &query)
    if err != nil {
        status = 1
        msg = "Database error"
    }
    if data[0]["isThere"] != "0" {
        status = 1
        msg = "User already exists"
    }else{
        /* insert user to database */
        /* get last usercode and add value to insert query */
        currentTime := time.Now()
        query = fmt.Sprintf("Insert into BU values(firstName, lastName, email, password, totalGroups, totalFiles, createdOn, isActive) values " +
                "('%s', '%s', '%s', '%s', '%d', '%d', '%s', '%s');",
                body["firstName"], body["lastName"], body["email"], body["password"], 0, 0, currentTime.Format("2006.01.02 15:04:05"), "TRUE");
        fmt.Println(query)
        _, err := db.CallDatabase(true, &query)
        if err != nil {
            status = 1
            msg = "Could not create User"
        }
    }

    var response Response
    response.Status = status
    response.Msg = msg
    response.Data = nil

    json.NewEncoder(res).Encode(response)
}

