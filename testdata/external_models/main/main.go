package main

import (
	"net/http"

	"github.com/swaggo/swag/example/basic/api"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @BasePath /v1
func main() {
	http.HandleFunc("/testapi/get-string-by-int/", api.GetStringByInt)
}
