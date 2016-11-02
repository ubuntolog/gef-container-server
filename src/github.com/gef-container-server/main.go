package main

import (
	"github.com/gef-container-server/api"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server starting")
	http.ListenAndServe(":8080", api.Handlers())
}