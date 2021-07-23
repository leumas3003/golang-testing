package config

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/labstack/gommon/log"
)

var c *AppConfig

const confFileName = "/Users/samuelmartel/WS/golang-testing/configs/golang-testing/conf.json"

// GetConfig returns the configuration
func GetConfig() (*AppConfig, error) {
	if c == nil {

		configFile := flag.String("conf", confFileName, "Fully qualified json configs file")
		flag.Parse()
		log.Infof("Loading configuration from [%v]", *configFile)

		file, err := os.Open(*configFile)
		if err != nil {
			log.Fatalf("Error opening file [%v]", err)
			flag.Usage()
			return nil, err

			// os.Exit(1)
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&c)
		if err != nil {
			log.Error(err)
			return nil, err
		}

	}

	return c, nil
}

// AppConfig contains the application configuration
type AppConfig struct {
	Database   database
	Messaging  messaging
	HttpServer httpserver
}

type database struct {
	User         string
	Password     string
	Host         string
	DbName       string
	Port         int
	Sslmode      string
	MaxOpenConns int
	MaxIdleConns int
}

type messaging struct {
	Url string
}

type httpserver struct {
	Addr            string
	ReadTimeout     int
	HandlerTimeout  int
	ExternalTimeout int
	WriteTimeout    int
}
