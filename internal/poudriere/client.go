package poudriere

import (
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

func (c *Client) CreatePorts(version string) error {
	cmd := exec.Command(c.executable, "ports", "-c", "-p", version)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (c *Client) UpdatePorts(version string) error {
	cmd := exec.Command(c.executable, "ports", "-u", "-p", version)
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

func (c *Client) BuildPackages(jail string, pkgs []string) error {
	args := append([]string{"bulk", "-j", jail}, pkgs...)
	cmd := exec.Command(c.executable, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
