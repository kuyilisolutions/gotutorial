package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// http.HandleFunc("web", processHandler)
	// http.Handle("/", http.FileServer(http.Dir("./static")))

	// fs := http.FileServer(http.Dir("../../web"))

	// // Handle "/" route to serve static files
	// http.Handle("/", fs)

	// fmt.Println("Server is running on :8080")

	fs := http.FileServer(http.Dir("web"))
	// Handle "/" route to serve static files
	http.Handle("/", fs)

	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		// Parse the incoming JSON data into a struct
		fmt.Println("Entering Go Backend")
		var hostDetails string
		fmt.Println(r.Body)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		hostDetails = string(body)
		ds := strings.Split(hostDetails, "=")[1]
		fmt.Println("Received data:", ds)
		if ds == "mani" {
			fmt.Println("Onna number somberi")
		}
		if ds == "kather" {
			fmt.Println("Lusuu pakki")
		}
		if ds == "bala" {
			fmt.Println("ultra level playboy")
		}

	})
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
