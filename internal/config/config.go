package config

type JailConfig struct {
	Version string `yaml:"version"`
	Arch    string `yaml:"arch"`
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
}
