package goperf

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ConfigFromEnv populates config from environment variables.
func ConfigFromEnv() (Config, error) {
	cfg := Config{}
	if v := os.Getenv("GOPERF_SERVICE_NAME"); v != "" {
		cfg.ServiceName = v
	}
	if v := os.Getenv("GOPERF_PATH_PREFIX"); v != "" {
		cfg.PathPrefix = v
	}
	if v := os.Getenv("GOPERF_PUBLIC_URL"); v != "" {
		cfg.PublicURL = v
	}
	if u := os.Getenv("GOPERF_USERNAME"); u != "" {
		cfg.Username = u
	}
	if p := os.Getenv("GOPERF_PASSWORD"); p != "" {
		cfg.Password = p
	}
	if d := os.Getenv("GOPERF_DB_DRIVER"); d != "" {
		cfg.StoreDriver = d
	}
	if d := os.Getenv("GOPERF_DB_DSN"); d != "" {
		cfg.StoreDSN = d
	}
	if v := os.Getenv("GOPERF_SMTP_PORT"); v != "" {
		if _, err := strconv.Atoi(v); err != nil {
			return cfg, fmt.Errorf("invalid GOPERF_SMTP_PORT: %w", err)
		}
	}
	return cfg, nil
}

// Validate verifies the monitor configuration.
func (c Config) Validate() error {
	if c.PathPrefix == "" {
		return fmt.Errorf("path prefix is required")
	}
	p := strings.TrimRight(c.PathPrefix, "/")
	if !strings.HasPrefix(p, "/") || p == "/" {
		return fmt.Errorf("path prefix must start with / and not be just /")
	}
	if !c.AllowNoAuth && c.Username == "" && c.Password == "" {
		return fmt.Errorf("authentication is required unless AllowNoAuth is true")
	}
	if c.StoreDriver != "" && c.StoreDriver != "sqlite" && c.StoreDriver != "postgres" {
		return fmt.Errorf("unsupported store driver: %s", c.StoreDriver)
	}
	return nil
}
