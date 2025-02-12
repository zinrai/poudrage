package config

import (
	"strings"
)

type JailConfig struct {
	Version string `yaml:"version"`
	Arch    string `yaml:"arch"`
}

type MakeConf struct {
	SystemDefined string
	UserDefined   string
}

type Environment struct {
	Jail     JailConfig `yaml:"jail"`
	MakeConf string     `yaml:"make.conf"`
}

type PackageOption struct {
	Name    string `yaml:"name"`
	Options string `yaml:"options"`
}

type Config struct {
	Environment Environment     `yaml:"environment"`
	Packages    []PackageOption `yaml:"packages"`
	MakeConf    MakeConf
}

func (c *Config) InitMakeConf() {
	c.MakeConf = MakeConf{
		UserDefined: c.Environment.MakeConf,
	}
}

func (m *MakeConf) String() string {
	var sb strings.Builder

	sb.WriteString("# System defined settings\n")
	sb.WriteString(m.SystemDefined)
	sb.WriteString("\n")

	sb.WriteString("# User defined settings\n")
	sb.WriteString(m.UserDefined)

	return sb.String()
}
