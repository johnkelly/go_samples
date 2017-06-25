package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HELLO WORLD")
}


