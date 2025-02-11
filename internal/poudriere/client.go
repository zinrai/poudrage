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

func (c *Client) CreateJail(name, version, arch string) error {
	cmd := exec.Command(c.executable, "jail", "-c", "-j", name, "-v", version, "-a", arch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (c *Client) UpdateJail(name string) error {
	cmd := exec.Command(c.executable, "jail", "-u", "-j", name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (c *Client) JailExists(name string) (bool, error) {
	cmd := exec.Command(c.executable, "jail", "-l", "-n")
	output, err := cmd.Output()
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
	cmd := exec.Command(c.executable, "ports", "-c", "-p", portsName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (c *Client) UpdatePorts(version string) error {
	portsName := FormatPortsName(version)
	cmd := exec.Command(c.executable, "ports", "-u", "-p", portsName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (c *Client) PortsExists(version string) (bool, error) {
	portsName := FormatPortsName(version)
	cmd := exec.Command(c.executable, "ports", "-l", "-n")
	output, err := cmd.Output()
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

func (c *Client) BuildPackages(jail string, pkgs []string) error {
	args := append([]string{"bulk", "-j", jail}, pkgs...)
	cmd := exec.Command(c.executable, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
