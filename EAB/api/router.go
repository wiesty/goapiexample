package api

import (
    "github.com/gorilla/mux"
    "EAB/api/example"
)

func SetupRouter() *mux.Router {
    router := mux.NewRouter()

    // Example Route
    router.HandleFunc("/api/example", example.HelloWorld).Methods("GET")

    return router
}
