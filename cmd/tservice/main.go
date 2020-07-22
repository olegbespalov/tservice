package main

import (
	"log"
	"net/http"

	"github.com/olegbespalov/tservice/pkg/config"
	"github.com/olegbespalov/tservice/pkg/handler"
	"github.com/olegbespalov/tservice/pkg/response"
)

func main() {
	cfg := config.NewService()
	responsesService := response.NewService(cfg)

	log.Println("TService is starting on port " + cfg.Port())
	log.Fatal(http.ListenAndServe(":"+cfg.Port(), handler.NewDefaultHandler(responsesService)))
}
