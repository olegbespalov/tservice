package config

import (
	"flag"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/olegbespalov/tservice/pkg/entity"
	"gopkg.in/yaml.v2"
)

var configFile string
var assetsPath string

func init() {
	flag.StringVar(&configFile, "config", "", "a path to the config file")
	flag.StringVar(&assetsPath, "assets", "", "an asset folder where we can find file responses")

	flag.Parse()

	if configFile == "" {
		log.Fatalln("You should provide a config file")
	}
}

type service struct {
	cfg        entity.Config
	assetsPath string
}

//NewService creates dummy config service
func NewService() UseCase {
	return service{
		cfg:        parseConfig(),
		assetsPath: assetsPath,
	}
}

func parseConfig() entity.Config {
	data, err := ioutil.ReadFile(filepath.Clean(configFile))
	if err != nil {
		log.Fatalf("config file isn't readable: %s\n", err.Error())
	}

	cfg := entity.Config{}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return cfg
}

func (s service) AssetPath() string {
	return s.assetsPath
}

func (s service) Config() entity.Config {
	return s.cfg
}

func (s service) ResponseDefinition() map[string]entity.ResponseDefinition {
	return s.cfg.ResponseDefinitions
}
