package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	html := `<html><body style='background-color: blue; display: flex; justify-content: center; align-items: center; height: 100vh;'><h1 style='color: white;'>It is working!</h1></body></html>`
	fmt.Fprintf(w, html)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handler)

	logger := log.New(os.Stdout, "devops-as-a-service: ", log.LstdFlags|log.Lshortfile)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: http.DefaultServeMux,
		ErrorLog: logger,
	}

	fmt.Printf("Server listening on port %s...\n", port)
	logger.Fatal(server.ListenAndServe())
}
