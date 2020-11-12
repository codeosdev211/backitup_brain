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
    router.HandleFunc("/files/list", routes.ListFiles)
    router.HandleFunc("/user/stats", routes.GetStats)
    router.HandleFunc("/files/download", routes.DownloadFile)
    router.HandleFunc("/groups/list", routes.GetGroups)
    router.HandleFunc("/groups/create", routes.CreateGroup)

    http.ListenAndServe(":8001", router)
}
