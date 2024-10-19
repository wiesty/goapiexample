package api

import (
    "net/http"
    "EAB/api/example"
)

func SetupRouter() {
    // Example Route
    http.HandleFunc("/api/example", example.HelloWorld)
}
