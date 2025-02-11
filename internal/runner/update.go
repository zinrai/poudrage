package runner

import (
	"fmt"

	"github.com/zinrai/poudrage/internal/config"
	"github.com/zinrai/poudrage/internal/poudriere"
)

func Update(cfg *config.Config) error {
	client := poudriere.NewClient()

	if err := client.UpdatePorts(cfg.Environment.Jail.Version); err != nil {
		return fmt.Errorf("failed to update ports: %w", err)
	}

	jailName := poudriere.FormatJailName(cfg.Environment.Jail.Version, cfg.Environment.Jail.Arch)
	if err := client.UpdateJail(jailName); err != nil {
		return fmt.Errorf("failed to update jail: %w", err)
	}

	return nil
}
