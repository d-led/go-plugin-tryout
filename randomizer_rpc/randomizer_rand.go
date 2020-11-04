package main

import (
	"github.com/d-led/go-plugin-tryout/interfaces"
	"github.com/d-led/go-plugin-tryout/randomizer"
	"github.com/hashicorp/go-plugin"
)

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "SAMPLE_PLUGIN",
	MagicCookieValue: "rand",
}

func main() {
	impl := randomizer.NewRand()
	var pluginMap = map[string]plugin.Plugin{
		// just one plugin here
		"randomizer": &interfaces.RandomizerPlugin{Impl: impl},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}
