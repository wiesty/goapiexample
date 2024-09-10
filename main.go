package main

import (
    "log"
    "net/http"
    "EAB/api"
)

func main() {
    router := api.SetupRouter()

    log.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", router)
    if err != nil {
        log.Fatal("Server failed to start:", err)
    }
}
