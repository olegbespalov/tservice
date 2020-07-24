package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/olegbespalov/tservice/pkg/config"
	"github.com/olegbespalov/tservice/pkg/handler"
	"github.com/olegbespalov/tservice/pkg/response"
)

var configFile string
var responsesPath string
var port string

func init() {
	flag.StringVar(&configFile, "config", "", "a path to the config file")
	flag.StringVar(&responsesPath, "responsePath", "", "a folder where we can find file responses")
	flag.StringVar(&port, "port", "", "a service port")

	flag.Parse()

	if configFile == "" {
		log.Fatalln("You should provide a config file")
	}

	if port == "" {
		log.Fatalln("You should provide a port")
	}
}

func main() {
	cfg := config.NewService(port, configFile, responsesPath)
	responsesService := response.NewService(cfg)

	log.Println("TService is starting on port " + cfg.Port())
	log.Fatal(http.ListenAndServe(":"+cfg.Port(), handler.NewDefaultHandler(responsesService)))
}
