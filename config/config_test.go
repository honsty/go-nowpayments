package config

import (
	"testing"
)

var validCfg = Credentials{
	Server:       "http://some.tld",
	Login:        "mylogin",
	Password:     "mypass",
	APIKey:       "key",
	IPNSecretKey: "key",
}

func TestLoad(t *testing.T) {
	emptyAPIKeyCfg := Credentials{Server: "http://some.tld", Login: "mylogin", Password: "mypass", IPNSecretKey: "key"}
	emptyLoginCfg := Credentials{Server: "http://some.tld", APIKey: "key", Password: "mypass", IPNSecretKey: "key"}
	emptyPasswordCfg := Credentials{Server: "http://some.tld", APIKey: "key", Login: "mylogin", IPNSecretKey: "key"}
	emptyServerCfg := Credentials{APIKey: "key", Login: "mylogin", Password: "mypass", IPNSecretKey: "key"}
	tests := []struct {
		name    string
		r       *Credentials
		wantErr bool
	}{
		{"nil reader", nil, true},
		{"bad config", &Credentials{}, true},
		{"valid config", &validCfg, false},
		{"empty API key", &emptyAPIKeyCfg, true},
		{"empty login", &emptyLoginCfg, true},
		{"empty password", &emptyPasswordCfg, true},
		{"empty server", &emptyServerCfg, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Load(tt.r); (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogin(t *testing.T) {
	Load(&validCfg)
	tests := []struct {
		name string
		want string
	}{
		{"login value", "mylogin"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Login(); got != tt.want {
				t.Errorf("Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassword(t *testing.T) {
	Load(&validCfg)
	tests := []struct {
		name string
		want string
	}{
		{"password value", "mypass"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Password(); got != tt.want {
				t.Errorf("Password() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPIKey(t *testing.T) {
	Load(&validCfg)
	tests := []struct {
		name string
		want string
	}{
		{"key value", "key"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := APIKey(); got != tt.want {
				t.Errorf("APIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer(t *testing.T) {
	Load(&validCfg)
	tests := []struct {
		name string
		want string
	}{
		{"server url", "http://some.tld"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Server(); got != tt.want {
				t.Errorf("Server() = %v, want %v", got, tt.want)
			}
		})
	}
}
