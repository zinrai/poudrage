package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zinrai/poudrage/internal/config"
	"github.com/zinrai/poudrage/internal/runner"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [setup|validate|build] -c build-env.yaml\n", os.Args[0])
		os.Exit(1)
	}

	command := os.Args[1]
	os.Args = append(os.Args[:1], os.Args[2:]...)

	configFile := flag.String("c", "build-env.yaml", "config file path")
	flag.Parse()

	cfg, err := config.Load(*configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}

	switch command {
	case "setup":
		if err := runner.Setup(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "setup failed: %v\n", err)
			os.Exit(1)
		}
	case "validate":
		if err := runner.Validate(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "validation failed: %v\n", err)
			os.Exit(1)
		}
	case "build":
		if err := runner.Build(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "build failed: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", command)
		fmt.Fprintf(os.Stderr, "Usage: %s [setup|validate|build] -c build-env.yaml\n", os.Args[0])
		os.Exit(1)
	}
}
