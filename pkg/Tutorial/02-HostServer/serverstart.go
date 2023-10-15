package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/web", processHandler)
    http.Handle("/", http.FileServer(http.Dir("./static")))
    
    fmt.Println("Server is running on :8080")
    http.ListenAndServe(":80", nil)
}

func processHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        r.ParseForm()
        data := r.FormValue("data")
        // Do something with the data, e.g., print it
        fmt.Println("Received data:", data)
    }
}
