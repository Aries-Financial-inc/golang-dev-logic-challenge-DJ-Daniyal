package main

import (
    "log"
    "net/http"
    "options-analysis/handlers"
)

func main() {
    http.HandleFunc("/analyze", handlers.AnalysisEndpoint)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
