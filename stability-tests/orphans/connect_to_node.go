package main

import (
	"fmt"
	"os"

	"github.com/Kalibriumnet/Kalibrium/infrastructure/config"
	"github.com/Kalibriumnet/Kalibrium/infrastructure/network/netadapter/standalone"
)

func connectToNode() *standalone.Routes {
	cfg := activeConfig()

	kalibriumdConfig := config.DefaultConfig()
	kalibriumdConfig.NetworkFlags = cfg.NetworkFlags

	minimalNetAdapter, err := standalone.NewMinimalNetAdapter(kalibriumdConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating minimalNetAdapter: %+v", err)
		os.Exit(1)
	}
	routes, err := minimalNetAdapter.Connect(cfg.NodeP2PAddress)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error connecting to node: %+v", err)
		os.Exit(1)
	}
	return routes
}
