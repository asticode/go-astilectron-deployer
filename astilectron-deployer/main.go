package main

import (
	"context"
	"flag"

	"github.com/asticode/go-astilectron-deployer"
	"github.com/asticode/go-astilog"
	"github.com/asticode/go-astitools/config"
	"github.com/asticode/go-astitools/flag"
	"github.com/pkg/errors"
	"github.com/asticode/go-astitools/os"
)

// Flags
var (
	configPath = flag.String("c", "", "the config path")
)

func main() {
	// Parse flags
	s := astiflag.Subcommand()
	flag.Parse()
	astilog.FlagInit()

	// Create configuration
	i, err := asticonfig.New(&Configuration{}, *configPath, &Configuration{})
	if err != nil {
		astilog.Fatal(errors.Wrap(err, "main: creating configuration failed"))
	}
	c := i.(*Configuration)

	// Create logger
	astilog.New(c.Logger)

	// Create deployer
	d := astideployer.New(c.Deployer)

	// Create context
	ctx, cancel := context.WithCancel(context.Background())

	// Handle signals
	go astios.HandleSignals(astios.ContextSignalsFunc(cancel))

	// Switch on subcommand
	switch s {
	case "private":
		if err = d.ServePrivate(ctx); err != nil {
			astilog.Fatal(errors.Wrap(err, "main: ServePrivate failed"))
		}
	}
}

type Configuration struct {
	Deployer astideployer.Configuration `toml:"deployer"`
	Logger   astilog.Configuration      `toml:"logger"`
}
