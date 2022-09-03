package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("start server")
	http.HandleFunc("/", handler)
	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Fprint(w, len(strings.Split(path, "/")))
}
