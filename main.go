package main

import (
    "log"
    "net/http"
    "EAB/api"
)

func main() {
    api.SetupRouter()
    
    log.Println("Server running on port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
