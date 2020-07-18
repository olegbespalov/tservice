package main

import (
	"log"
	"net/http"

	"github.com/olegbespalov/tservice/pkg/config"
	"github.com/olegbespalov/tservice/pkg/handler"
)

var port = "8080"

func main() {

	configService := config.NewService()

	log.Println("TService is starting on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, handler.NewDefaultHandler(configService)))
}
