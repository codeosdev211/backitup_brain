package main

import (
    http "net/http"
    "encoding/json"
    mux "github.com/gorilla/mux"
    db "./db"
    _ "github.com/go-sql-driver/mysql"
)

type Response struct {
    Status int `json:"status"`
    Msg string `json:"msg"`
    Data []map[string]interface{} `json:"data"`
}

/* main function allows routing requests */
func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/home", homePage)

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
    var status int = 0
    var msg string = "none"
    setupCORS(&res)
    var response Response
    response.Status = status
    response.Msg = msg
    json.NewEncoder(res).Encode(response)
}

/* signin */
func signInUser(res http.ResponseWriter, req *http.Request) {
    var request map[string]interface{}
    var msg string = "none"
    var status int8 = 0

    err := json.Decoder(req.Body).Decode(&request)
    if err != nil {
        status = 1
        msg = "Invalid request body"
    }
    body := request["values"][0]
    var query string = "select count(*) as isThere from BU where email='" + body["email"] + "' and password='"+ body["password"] +"';"

    data, err := db.CallDatabase(1, &query)
    if err != nil {
        status = 1
        msg = "Database error"
    }

    if data[0]["isThere"] == 0 {
        status = 1
        msg = "Invalid Email or Password"
    }else{
        query = "select * from BU where email='" + body["email"] + "' and password='"+ body["password"] +"';" 
        data, err := db.CallDatabase(1, &query)
        if err != nil {
            status = 1
            msg = "Could not get user data"
        }
    }
    var response Response
    response.Status = status
    response.Msg = msg
    responsg.Data = data

    json.NewEncoder(res).Encode(response)
}


func signUpUser(res http.ResponseWriter, req *http.Request) {
    var request map[string]interface{}
    var msg string = "none"
    var status int8 = 0

    err := json.Decoder(req.Body).Decode(&request)
    if err != nil {
        status = 1
        msg = "Invalid request body"
    }
    body := request["values"][0]
    var query string = "select count(*) as isThere from BU where email='" + body["email"] + "' and password='"+ body["password"] +"';"

    data, err := db.CallDatabase(1, &query)
    if err != nil {
        status = 1
        msg = "Database error"
    }

    if data[0]["isThere"] != 0 {
        status = 1
        msg = "User already exists"
    }else{
        /* insert user to database */

    }

    var response Response
    response.Status = status
    response.Msg = msg
    responsg.Data = nil

    json.NewEncoder(res).Encode(response)
}

