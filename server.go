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
    data, err := db.CallDatabase(1, "select * from AUser")
    if err != nil {
        status = 1
        msg = "No records"
    }
    setupCORS(&res)
    var response Response
    response.Status = status
    response.Msg = msg
    response.Data = data

    json.NewEncoder(res).Encode(response)
}



