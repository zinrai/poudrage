package runner

import (
	"github.com/zinrai/poudrage/internal/config"
	"github.com/zinrai/poudrage/internal/poudriere"
)

func Build(cfg *config.Config, configFile string) error {
	client := poudriere.NewClient()
	setname := config.ExtractSetName(configFile)

	var pkgs []string
	for _, pkg := range cfg.Packages {
		pkgs = append(pkgs, pkg.Name)
	}

	jailName := poudriere.FormatJailName(cfg.Environment.Jail.Version, cfg.Environment.Jail.Arch)
	return client.BuildPackages(jailName, cfg.Environment.Jail.Version, setname, pkgs)
}
