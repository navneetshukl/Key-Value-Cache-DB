package main

import (
	"key-value-db/client"
	"net/http"
)

func main() {

	http.HandleFunc("/", client.Client)
	http.ListenAndServe(":8081", nil)
}
