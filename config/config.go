package config

import (
	"errors"
	"net/url"

	"github.com/rotisserie/eris"
)

type Credentials struct {
	APIKey       string `json:"apiKey"`
	IPNSecretKey string `json:"ipnSecretKey"`
	Login        string `json:"login"`
	Password     string `json:"password"`
	Server       string `json:"server"`
}

var conf *Credentials

func configErr(err error) error {
	return eris.Wrap(err, "config")
}

// Load parses a JSON file to get the required credentials to operate NOWPayment's API.
func Load(c *Credentials) error {
	if c == nil {
		return configErr(errors.New("nil reader"))
	}
	if conf == nil {
		conf = c
	} else {
		conf.APIKey = c.APIKey
		conf.Server = c.Server
		conf.IPNSecretKey = c.IPNSecretKey
		conf.Login = c.Login
		conf.Password = c.Password
	}

	// Sanity checks.
	if conf.APIKey == "" {
		return configErr(errors.New("API key is missing"))
	}
	if conf.IPNSecretKey == "" {
		return configErr(errors.New("IPN secret key is missing"))
	}

	if conf.Server == "" {
		return configErr(errors.New("server URL missing"))
	}

	_, err := url.Parse(conf.Server)
	if err != nil {
		return configErr(errors.New("server URL parsing"))
	}

	return nil
}

// Login returns the email address to use with the API.
func Login() string {
	return conf.Login
}

// Password returns the related password to use.
func Password() string {
	return conf.Password
}

// APIKey is the API key to use.
func APIKey() string {
	return conf.APIKey
}

// IPNSecretKey returns the related IPN secret key to use.
func IPNSecretKey() string {
	return conf.IPNSecretKey
}

// Server returns URL to connect to the API service.
func Server() string {
	return conf.Server
}
