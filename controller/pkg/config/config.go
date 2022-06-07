package config

import (
	"os"
	"path/filepath"

	"cornerstone_verifier/pkg/log"

	"github.com/joho/godotenv"
)

// Config stores all env vars.
type Config struct {
}

// GetConfig returns the config.
func GetConfig() *Config {
	config := &Config{}

	envFilePath, _ := filepath.Abs(".env")

	if err := godotenv.Load(envFilePath); err != nil {
		log.ServerWarning.Print("No .env file found, using env vars from os.")
	}

	return config
}

// getEnv loads env vars from .env file or os.
func getEnv(key string) string {
	return os.Getenv(key)
}

// GetAcapyURL returns the Acapy URL.
func (c *Config) GetAcapyURL() string {
	return getEnv("ACAPY_URL")
}

// GetCredDefID returns the Cred Def ID.
func (c *Config) GetCredDefID() string {
	return getEnv("CRED_DEF_ID")
}

// GetClientURL returns the client URL.
func (c *Config) GetClientURL() string {
	return getEnv("CLIENT_URL")
}

// GetServerAddress returns the server address.
func (c *Config) GetServerAddress() string {
	return getEnv("SERVER_ADDRESS")
}

// GetServerBaseURL returns the api base url.
func (c *Config) GetAPIBaseURL() string {
	return getEnv("API_BASE_URL")
}
