package config

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

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
	modified   time.Time

	mu sync.Mutex
}

//NewService creates dummy config service
func NewService() UseCase {
	return &service{
		cfg:        parseConfig(),
		modified:   configModified(),
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

func configModified() time.Time {
	info, err := os.Stat(filepath.Clean(configFile))
	if err != nil {
		log.Fatalf("config file isn't readable: %s\n", err.Error())
	}

	return info.ModTime()
}

func (s service) AssetPath() string {
	return s.assetsPath
}

func (s *service) Config() entity.Config {
	s.mu.Lock()
	defer s.mu.Unlock()

	current := configModified()
	if s.modified != current {
		s.cfg = parseConfig()
		s.modified = current
	}

	return s.cfg
}

func (s service) ResponseDefinition() map[string]entity.ResponseDefinition {
	return s.Config().ResponseDefinitions
}
