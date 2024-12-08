package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AzureOpenai LlmConfig `split_words:"true"`
	Openai      LlmConfig `split_words:"true"`
	Llama31     LlmConfig `split_words:"true"`
	Claude      LlmConfig `split_words:"true"`
	Perplexity  LlmConfig `split_words:"true"`
	Storage     StorageProvider
	ServiceBus  ServiceBusConfig `split_words:"true"`
}

type LlmConfig struct {
	Endpoint   string `required:"true"`
	ApiKey     string `required:"true"`
	ModelName  string
	ApiVersion string
}

type StorageProvider struct {
	Provider string       `required:"true"`
	Config   CloudStorage `split_words:"true"`
}

type CloudStorage struct {
	Endpoint    string `required:"true"`
	ApiKey      string `required:"true"`
	BucketName  string `required:"true"`
	AccountName string `required:"true"`
}

type ServiceBusConfig struct {
	ConnectionString string `split_words:"true"`
}

func Init() (*Config, error) {
	if isInContainer() {
		fmt.Println("Running in container, not loading .env")
	} else {
		fmt.Println("Loading environment variables from .env")
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading environment variables:", err)
			return nil, err
		}
	}

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Printf("Failed to process environment variables: %v", err)
		return nil, err
	}

	return &cfg, nil
}

func isInContainer() bool {
	// Check if the environment variable indicating a container is set
	return os.Getenv("RUNNING_IN_CONTAINER") == "true"
}
