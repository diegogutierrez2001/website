package main

import (
	"net/http"
	"log"
	"fmt"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello! I'm %s.\n", r.RemoteAddr)
	fmt.Fprintf(w, "Hello! I'm Diego.")
}

func main() {
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
