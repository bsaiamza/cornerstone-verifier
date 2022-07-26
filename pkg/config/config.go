package config

import (
	"os"
)

type Config struct {
}

func LoadConfig() *Config {
	return &Config{}
}

func getEnvVarByName(name string) string {
	return os.Getenv(name)
}

func (c *Config) GetAcapyURL() string {
	return getEnvVarByName("ACAPY_URL")
}

func (c *Config) GetCornerstoneCredDefID() string {
	return getEnvVarByName("CORNERSTONE_CRED_DEF_ID")
}

func (c *Config) GetAddressCredDefID() string {
	return getEnvVarByName("ADDRESS_CRED_DEF_ID")
}

func (c *Config) GetVaccineCredDefID() string {
	return getEnvVarByName("VACCINE_CRED_DEF_ID")
}

func (c *Config) GetServerAddress() string {
	return getEnvVarByName("SERVER_ADDRESS")
}
