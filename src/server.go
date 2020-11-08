package main

import (
    http "net/http"
    mux "github.com/gorilla/mux"
    routes "./routes"
)

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/user/login", routes.SignInUser)
    router.HandleFunc("/user/register", routes.SignUpUser)
    router.HandleFunc("/files/add", routes.AddFiles)
//    router.HandleFunc("/files/list", routes.ListFiles)

    http.ListenAndServe(":8001", router)
}
