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

func (c *Config) GetCornerstoneSchemaID() string {
	return getEnvVarByName("CORNERSTONE_SCHEMA_ID")
}

func (c *Config) GetAddressSchemaID() string {
	return getEnvVarByName("ADDRESS_SCHEMA_ID")
}

func (c *Config) GetContactableCredDefID() string {
	return getEnvVarByName("CONTACTABLE_CRED_DEF_ID")
}

func (c *Config) GetServerAddress() string {
	return getEnvVarByName("SERVER_ADDRESS")
}

func (c *Config) GetEmailUsername() string {
	return getEnvVarByName("EMAIL_USERNAME")
}

func (c *Config) GetEmailPassword() string {
	return getEnvVarByName("EMAIL_PASSWORD")
}

func (c *Config) GetSmtpServer() string {
	return getEnvVarByName("EMAIL_SMTP_SERVER")
}

func (c *Config) GetSmtpPort() string {
	return getEnvVarByName("EMAIL_SMTP_PORT")
}

func (c *Config) GetTxnCounterAPI() string {
	return getEnvVarByName("TXN_COUNTER_API")
}

func (c *Config) GetTxnCounterSwitch() string {
	return getEnvVarByName("TXN_COUNTER_SWITCH")
}

func (c *Config) GetPValueYear() string {
	return getEnvVarByName("P_VALUE_YEAR")
}
