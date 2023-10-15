package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", fs)

	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		// Parse the incoming form values
		name := r.FormValue("name")
		data := r.FormValue("data")

		// Process the data based on the selected name
		response := processFormData(name, data)

		// Send a response back to the webpage
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	})

	http.ListenAndServe(":8080", nil)
}

func processFormData(name, data string) string {
	// Process the data based on the selected name
	var response string
	switch name {
	case "mani":
		response = "Mani Prakash: " + data + "Onna number somberi" + "Developer"
	case "aravind":
		response = "Aravind: " + data + "Vayadi" + "Developer"
	case "kather":
		response = "Katheravan: " + data + "Loosu pakki" + "Developer"
	case "nitty":
		response = "Nitty: " + data + "solldraku onnu ilaa" + "Developer"
	default:
		response = "Unknown name"
	}
	return response
}
