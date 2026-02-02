package configs

import (
	"fmt"

	"github.com/caarlos0/env/v10"
)

// Config holds the configuration for the application.
type Config struct {
	Port        int    `env:"LSE_PORT" envDefault:"8080"`
	Debug       bool   `env:"LSE_DEBUG" envDefault:"false"`
	JSONLog     bool   `env:"LSE_JSON_LOG" envDefault:"false"`
	Interval    string `env:"LSE_INTERVAL" envDefault:"10s"`
	NodeName    string `env:"LSE_NODE_NAME" envDefault:""`
	CertFile    string `env:"LSE_CERT_FILE" envDefault:"/var/lib/kubelet/pki/kubelet-client-current.pem"`
	KeyFile     string `env:"LSE_KEY_FILE" envDefault:"/var/lib/kubelet/pki/kubelet-client-current.pem"`
	K8SLocalAPI string `env:"LSE_K8S_LOCAL_API" envDefault:""`
}

// LoadConfig loads the configuration from environment variables using the caarlos0/env package.
func LoadConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse environment variables: %w", err)
	}

	return &cfg, nil
}
