package config

import (
	"log"
	"path/filepath"

	"github.com/jkrus/kit/config"
	"github.com/jkrus/kit/files"
	"github.com/pkg/errors"
)

type (
	// Config represents the main app's configuration.
	Config struct {
		Host string `yaml:"host"`
		HTTP HTTP   `yaml:"http"`
		GRPC GRPC   `yaml:"grpc"`
		DB   DB     `yaml:"db"`
	}
)

// NewConfig constructs an empty configuration.
func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Init() error {
	filePath := filepath.Join(files.OsAppRootPath(AppRootPathName, AppName, AppUsage, AppVersion), yamlFileName)
	if files.IsFileExist(filePath) {
		log.Println("Read data from config file in path:", filePath)
		if err := files.ReadFromYamlFile(filePath, c); err != nil {
			return errors.Wrap(err, "Init: read config file filed")
		}
	} else {
		log.Println("Create default config file in path:", filePath)
	}

	c.setDefault()
	if err := files.MakeDirs(filePath); err != nil {
		return errors.Wrap(err, "Init: can not create dirs")
	}
	if err := files.WriteToYamlFile(filePath, c); err != nil {
		return errors.Wrap(err, "Init: create config file filed")
	}

	return nil
}

func (c *Config) Load() error {
	err := config.Load(AppRootPathName, AppName, AppUsage, AppVersion, yamlFileName, c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) setDefault() {
	c.Host = "127.0.0.1"

	c.HTTP.Port = 8080

	c.GRPC.Port = 8081

	c.DB.User = "root"
	c.DB.Pass = "secret"
	c.DB.Host = "localhost"
	c.DB.Port = 3306
	c.DB.Name = "blogbd"
	c.DB.Code = "utf8"
}
