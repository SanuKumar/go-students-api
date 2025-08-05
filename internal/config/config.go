package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

// env-required:"prod"
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"` // capital so that we can use it in templates --> struct tags
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config { // if function like must, do not return err, throw fatal error
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file") // passed as flag --> go run cmd/students-api/main.go -config-path abc
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) { // wrote two line in one
		log.Fatalf("config file does not exist: %s", configPath) // we are formatting here, so use Fatalf
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("can not read config file: %s", err.Error())
	}

	return &cfg
}
