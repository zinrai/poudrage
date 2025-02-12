package poudriere

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Client struct {
	executable string
}

func NewClient() *Client {
	return &Client{
		executable: "poudriere",
	}
}

func FormatPortsName(version string) string {
	name := strings.ReplaceAll(version, ".", "_")
	return strings.ReplaceAll(name, "-", "_")
}

func FormatJailName(version, arch string) string {
	name := version + "-" + arch
	return strings.ReplaceAll(name, ".", "_")
}

func (c *Client) runCommand(args ...string) error {
	fmt.Printf("Executing: %s %s\n", c.executable, strings.Join(args, " "))
	cmd := exec.Command(c.executable, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (c *Client) runCommandOutput(args ...string) ([]byte, error) {
	fmt.Printf("Executing: %s %s\n", c.executable, strings.Join(args, " "))
	return exec.Command(c.executable, args...).Output()
}

func (c *Client) SetupDistfilesCache() error {
	data, err := os.ReadFile("/usr/local/etc/poudriere.conf")
	if err != nil {
		return fmt.Errorf("failed to read poudriere.conf: %w", err)
	}

	lines := strings.Split(string(data), "\n")
	var distfilesDir string
	for _, line := range lines {
		if strings.HasPrefix(line, "DISTFILES_CACHE=") {
			distfilesDir = strings.TrimPrefix(line, "DISTFILES_CACHE=")
			break
		}
	}

	if distfilesDir == "" {
		return fmt.Errorf("DISTFILES_CACHE not found in poudriere.conf")
	}

	fmt.Printf("Creating distfiles directory: %s\n", distfilesDir)
	if err := os.MkdirAll(distfilesDir, 0755); err != nil {
		return fmt.Errorf("failed to create distfiles directory: %w", err)
	}

	return nil
}

func (c *Client) WriteMakeConf(jail, version, setname, makeconf string) error {
	portsName := FormatPortsName(version)
	makeconfPath := fmt.Sprintf("/usr/local/etc/poudriere.d/%s-%s-%s-make.conf",
		jail, portsName, setname)

	fmt.Printf("Writing make.conf to: %s\n", makeconfPath)
	if err := os.WriteFile(makeconfPath, []byte(makeconf), 0644); err != nil {
		return fmt.Errorf("failed to write make.conf: %w", err)
	}

	return nil
}

func (c *Client) WriteOptions(jail, version, setname, options string) error {
	portsName := FormatPortsName(version)
	optionsPath := fmt.Sprintf("/usr/local/etc/poudriere.d/%s-%s-%s-options",
		jail, portsName, setname)

	fmt.Printf("Writing options to: %s\n", optionsPath)
	if err := os.WriteFile(optionsPath, []byte(options), 0644); err != nil {
		return fmt.Errorf("failed to write options file: %w", err)
	}

	return nil
}

func (c *Client) CreateJail(name, version, arch string) error {
	return c.runCommand("jail", "-c", "-j", name, "-v", version, "-a", arch)
}

func (c *Client) UpdateJail(name string) error {
	return c.runCommand("jail", "-u", "-j", name)
}

func (c *Client) JailExists(name string) (bool, error) {
	output, err := c.runCommandOutput("jail", "-l", "-n")
	if err != nil {
		return false, fmt.Errorf("failed to list jails: %w", err)
	}

	jails := strings.Split(string(output), "\n")
	for _, jail := range jails {
		if strings.TrimSpace(jail) == name {
			return true, nil
		}
	}
	return false, nil
}

func (c *Client) CreatePorts(version string) error {
	portsName := FormatPortsName(version)
	return c.runCommand("ports", "-c", "-p", portsName)
}

func (c *Client) UpdatePorts(version string) error {
	portsName := FormatPortsName(version)
	return c.runCommand("ports", "-u", "-p", portsName)
}

func (c *Client) PortsExists(version string) (bool, error) {
	portsName := FormatPortsName(version)
	output, err := c.runCommandOutput("ports", "-l", "-n")
	if err != nil {
		return false, fmt.Errorf("failed to list ports: %w", err)
	}

	trees := strings.Split(string(output), "\n")
	for _, tree := range trees {
		if strings.TrimSpace(tree) == portsName {
			return true, nil
		}
	}
	return false, nil
}

func (c *Client) BuildPackages(jail string, version string, setname string, pkgs []string) error {
	portsName := FormatPortsName(version)
	args := append([]string{"bulk", "-c", "-j", jail, "-p", portsName, "-z", setname}, pkgs...)
	return c.runCommand(args...)
}
