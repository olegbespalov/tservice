package config

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/olegbespalov/tservice/pkg/entity"
	"gopkg.in/yaml.v2"
)

type service struct {
	port          string
	configFile    string
	cfg           entity.Config
	responsesPath string
	modified      time.Time

	mu sync.Mutex
}

//NewService creates dummy config service
func NewService(port, configFile, responsesPath string) UseCase {
	return &service{
		configFile:    configFile,
		port:          port,
		cfg:           parseConfig(configFile),
		modified:      configModified(configFile),
		responsesPath: responsesPath,
	}
}

func parseConfig(configFile string) entity.Config {
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

func configModified(configFile string) time.Time {
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

	current := configModified(s.configFile)
	if s.modified != current {
		s.cfg = parseConfig(s.configFile)
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
