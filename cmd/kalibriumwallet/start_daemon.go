package main

import "github.com/Kalibriumnet/Kalibrium/cmd/kalibriumwallet/daemon/server"

func startDaemon(conf *startDaemonConfig) error {
	return server.Start(conf.NetParams(), conf.Listen, conf.RPCServer, conf.KeysFile, conf.Profile, conf.Timeout)
}
