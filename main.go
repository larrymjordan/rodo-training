package main

import (
	"challenge2/ipScanner"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ip/details", ipScanner.StartIPScan)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
