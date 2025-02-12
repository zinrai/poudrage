package config

import (
	"fmt"
	"strings"
)

type JailConfig struct {
	Version string `yaml:"version"`
	Arch    string `yaml:"arch"`
}

type MakeConf string

type Options struct {
	Packages []PackageOption
}

type Environment struct {
	Jail     JailConfig `yaml:"jail"`
	MakeConf MakeConf   `yaml:"make.conf"`
}

type PackageOption struct {
	Name    string `yaml:"name"`
	Options string `yaml:"options"`
}

type Config struct {
	Environment Environment     `yaml:"environment"`
	Packages    []PackageOption `yaml:"packages"`
	Options     Options
}

func (c *Config) InitOptions() {
	c.Options = Options{
		Packages: c.Packages,
	}
}

func (m MakeConf) String() string {
	return string(m)
}

func (o *Options) String() string {
	var sb strings.Builder

	for i, pkg := range o.Packages {
		sb.WriteString(fmt.Sprintf("# %s\n", pkg.Name))
		sb.WriteString(pkg.Options)
		if i < len(o.Packages)-1 {
			sb.WriteString("\n\n")
		}
	}

	return sb.String()
}
