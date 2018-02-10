package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/HealthCheck", helloRequest)
	http.HandleFunc("/", getRequest)

	// start web server
	log.Println("Listening on http://localhost:5000/")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

// basic handler for /HealthCheck request
func helloRequest(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "It works! In future I will do more things :)")
	return
}

// this function simply serves up the file requested, or
// an index list if a directory is requested
func getRequest(w http.ResponseWriter, r *http.Request) {
	file_requested := "./" + r.URL.Path
	http.ServeFile(w, r, file_requested)
	return
}
