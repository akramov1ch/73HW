package main

import (
    "log"
    "net/http"
    "73HW/producer/router"
)

func main() {
    r := router.SetupRouter()
    log.Fatal(http.ListenAndServe(":8080", r))
}
