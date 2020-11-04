package main

import (
	"fmt"
	"os/exec"
	"plugin"

	"github.com/d-led/go-plugin-tryout/interfaces"
	hplugin "github.com/hashicorp/go-plugin"
)

func main() {
	fmt.Println("==== hashicorp/go-plugin demo ====")
	hashicorpPluginExample()
	fmt.Println("==== native go plugin demo =======")
	nativePluginExample()
	fmt.Println("==================================")
}

func hashicorpPluginExample() {
	const pluginCommand = "./randomizer_rpc_plugin"

	var handshakeConfig = hplugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "SAMPLE_PLUGIN",
		MagicCookieValue: "rand",
	}

	var pluginMap = map[string]hplugin.Plugin{
		// just one plugin here
		"randomizer": &interfaces.RandomizerPlugin{},
	}

	client := hplugin.NewClient(&hplugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command(pluginCommand),
	})
	defer client.Kill()

	// Connect to the plugin via RPC
	rpcClient, err := client.Client()
	if err != nil {
		fmt.Printf("Could not connect to the plugin: %v\n", err)
		return
	}

	// Get the plugin instance
	raw, err := rpcClient.Dispense("randomizer")
	if err != nil {
		fmt.Printf("Could not call the plugin: %v\n", err)
		return
	}
	r := raw.(interfaces.Randomizer)

	fmt.Println("some random number from a plugin:", r.Get())
}

func nativePluginExample() {
	myPlugin, err := plugin.Open("./randomizer_plugin.so")
	if err != nil {
		fmt.Printf("Could not open the native plugin: %v\n", err)
		return
	}

	raw, err := myPlugin.Lookup("InjectRand")
	if err != nil {
		fmt.Printf("Could not find entry point: %v\n", err)
		return
	}

	injector, ok := raw.(func(*interfaces.Randomizer))
	if !ok {
		fmt.Printf("Could not cast the type to the Randomizer interface: %v\n", err)
		return
	}

	var r interfaces.Randomizer
	// a new instance will be injected here
	injector(&r)

	fmt.Println("some random number from a plugin:", r.Get())
}
