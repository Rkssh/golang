package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `<html><body style='background-color: blue; display: flex; justify-content: center; align-items: center; height: 100vh;'><h1 style='color: white;'>It is working!</h1></body></html>`
		fmt.Fprintf(w, html)
	})

	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
