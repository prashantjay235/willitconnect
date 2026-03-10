package main

import (
	"log"
	"net/http"

	"github.com/prashantjay235/willitconnect/handlers"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./web")))

	http.HandleFunc("/dns", handlers.DNSCheck)
	http.HandleFunc("/tcp", handlers.TCPCheck)
	http.HandleFunc("/http", handlers.HTTPCheck)
	http.HandleFunc("/ping", handlers.PingCheck)
	http.HandleFunc("/traceroute", handlers.TraceRoute)

	log.Println("Server running on :8080")

	http.ListenAndServe(":8080", nil)

}
