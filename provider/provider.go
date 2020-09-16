package provider

import (
	"log"

	"github.com/ybbus/jsonrpc"
)

// Provider struct
type Provider struct {
	config *Config
	client jsonrpc.RPCClient
}

// New initiates the Provider struct
func New(c *Config) (p *Provider) {
	p = new(Provider)
	p.config = c

	p.init()

	return p
}

// Send executes the json-rpc request
func (p *Provider) Send(method string, args ...interface{}) *jsonrpc.RPCResponse {
	response, err := p.client.Call(method, args...)

	if err != nil {
		log.Fatal(err)
	}

	return response
}

func (p *Provider) init() {
	p.client = jsonrpc.NewClient(p.config.URL)
}
