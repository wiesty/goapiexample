package main

import (
    "log"
    "net/http"
    "EAB/api"
)

func main() {
    // Router initialisieren
    router := api.SetupRouter()

    // Server auf Port 8080 starten
    log.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", router)
    if err != nil {
        log.Fatal("Server failed to start:", err)
    }
}
