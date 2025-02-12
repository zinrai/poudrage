package runner

import (
	"fmt"

	"github.com/zinrai/poudrage/internal/config"
	"github.com/zinrai/poudrage/internal/poudriere"
)

func Setup(cfg *config.Config) error {
	client := poudriere.NewClient()

	if err := client.SetupDistfilesCache(); err != nil {
		return err
	}

	jailName := poudriere.FormatJailName(cfg.Environment.Jail.Version, cfg.Environment.Jail.Arch)

	if err := client.WriteMakeConf(jailName, cfg.Environment.Jail.Version, cfg.Environment.MakeConf.String()); err != nil {
		return err
	}

	if err := client.WriteOptions(jailName, cfg.Environment.Jail.Version, cfg.Options.String()); err != nil {
		return err
	}

	exists, err := client.JailExists(jailName)
	if err != nil {
		return err
	}

	if !exists {
		if err := client.CreateJail(jailName, cfg.Environment.Jail.Version, cfg.Environment.Jail.Arch); err != nil {
			return fmt.Errorf("failed to create jail: %w", err)
		}
		fmt.Printf("Created jail: %s\n", jailName)
	} else {
		fmt.Printf("Jail already exists: %s\n", jailName)
	}

	exists, err = client.PortsExists(cfg.Environment.Jail.Version)
	if err != nil {
		return err
	}

	if !exists {
		if err := client.CreatePorts(cfg.Environment.Jail.Version); err != nil {
			return fmt.Errorf("failed to create ports: %w", err)
		}
		fmt.Printf("Created ports: %s\n", cfg.Environment.Jail.Version)
	} else {
		fmt.Printf("Ports already exists: %s\n", cfg.Environment.Jail.Version)
	}

	return nil
}
