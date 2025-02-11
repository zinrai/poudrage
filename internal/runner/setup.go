package runner

import (
	"fmt"

	"github.com/zinrai/poudrage/internal/config"
	"github.com/zinrai/poudrage/internal/poudriere"
)

func Setup(cfg *config.Config) error {
	client := poudriere.NewClient()

	jailName := poudriere.FormatJailName(cfg.Environment.Jail.Version, cfg.Environment.Jail.Arch)

	if err := client.CreateJail(jailName, cfg.Environment.Jail.Version, cfg.Environment.Jail.Arch); err != nil {
		return fmt.Errorf("failed to create jail: %w", err)
	}

	if err := client.CreatePorts(cfg.Environment.Jail.Version); err != nil {
		return fmt.Errorf("failed to create ports: %w", err)
	}

	return nil
}
