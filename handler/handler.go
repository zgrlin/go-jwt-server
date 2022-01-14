package main

import (
	"fmt"
	"log"
	"net/http"
)

func web(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "wallet_id")
}

func authorized(endpoint func(http.ResponseWriter, *httpRequest)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
		} else {
			fmt.Fprintf(w, "Not Allowed")
		}
	})
}

func handleRequests() {
	http.Handle("/web", authorized(web))
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func main() {
	fmt.Println("Web Server Run..")
	handleRequests()
}
