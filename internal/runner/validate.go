package runner

import (
	"fmt"

	"github.com/zinrai/poudrage/internal/config"
)

func Validate(cfg *config.Config) error {
	if cfg.Environment.Jail.Version == "" {
		return fmt.Errorf("jail version is required")
	}
	if cfg.Environment.Jail.Arch == "" {
		return fmt.Errorf("jail arch is required")
	}

	return nil
}
