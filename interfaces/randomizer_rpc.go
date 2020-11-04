package interfaces

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

// RandomizerRPC client
type RandomizerRPC struct {
	client *rpc.Client
}

// Get via rpc
func (g *RandomizerRPC) Get() int {
	var res int
	err := g.client.Call("Plugin.Get", new(interface{}), &res)
	if err != nil {
		panic(err)
	}

	return res
}

// RandomizerRPCServer state
type RandomizerRPCServer struct {
	Impl Randomizer
}

// Get impl wrapper
func (s *RandomizerRPCServer) Get(args interface{}, res *int) error {
	*res = s.Impl.Get()
	return nil
}

// RandomizerPlugin boilerplate
type RandomizerPlugin struct {
	Impl Randomizer
}

// Server constructor
func (p *RandomizerPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &RandomizerRPCServer{Impl: p.Impl}, nil
}

// Client constructor
func (RandomizerPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &RandomizerRPC{client: c}, nil
}
