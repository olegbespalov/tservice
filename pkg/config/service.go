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

type service struct {
	port          string
	cfg           entity.Config
	responsesPath string
	modified      time.Time

	mu sync.Mutex
}

//NewService creates dummy config service
func NewService() UseCase {
	return &service{
		port:          port,
		cfg:           parseConfig(),
		modified:      configModified(),
		responsesPath: responsesPath,
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

//ResponsesPath path to the responses
func (s *service) ResponsesPath() string {
	return s.responsesPath
}

//config service config
func (s *service) config() entity.Config {
	s.mu.Lock()
	defer s.mu.Unlock()

	current := configModified()
	if s.modified != current {
		s.cfg = parseConfig()
		s.modified = current
	}

	return s.cfg
}

//Rules service rules
func (s *service) Rules() map[string]entity.Rule {
	return s.config().Rules
}

//Port service port
func (s *service) Port() string {
	return s.port
}

//DefaultDefinition default definition
func (s *service) DefaultDefinition() *entity.Definition {
	if v, ok := s.config().Default["response"]; ok {
		return &v
	}

	return nil
}

//DefaultError return default error if defined
func (s *service) DefaultError() *entity.Definition {
	if v, ok := s.config().Default["error"]; ok {
		return &v
	}

	return nil
}
