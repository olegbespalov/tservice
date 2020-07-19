package main

import (
	"log"
	"net/http"

	"github.com/olegbespalov/tservice/pkg/config"
	"github.com/olegbespalov/tservice/pkg/handler"
	"github.com/olegbespalov/tservice/pkg/response"
)

var port = "8080"

func main() {
	responsesService := response.NewService(config.NewService())

	log.Println("TService is starting on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, handler.NewDefaultHandler(responsesService)))
}
