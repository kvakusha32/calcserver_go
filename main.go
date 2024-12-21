package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/calculate", CalculateHandler)
	http.ListenAndServe(":8080", nil)
}
