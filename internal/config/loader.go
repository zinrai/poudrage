package config

import (
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	cfg.InitOptions()

	return &cfg, nil
}

func ExtractSetName(path string) string {
	base := filepath.Base(path)
	noExt := strings.TrimSuffix(base, filepath.Ext(base))
	parts := strings.FieldsFunc(noExt, func(r rune) bool {
		return r == '_' || r == '-'
	})
	if len(parts) > 0 {
		return parts[0]
	}
	return noExt
}
