package provider

import (
	"log"

	"github.com/ybbus/jsonrpc"
)

type Provider struct {
	config *Config
	client jsonrpc.RPCClient
}

func New(c *Config) (p *Provider) {
	p = new(Provider)
	p.config = c

	p.init()
	
	return p
}

func (p *Provider) init() {
	p.client = jsonrpc.NewClient(p.config.URL)
}

func (p *Provider) Send(method string, args ...interface{}) *jsonrpc.RPCResponse {
	response, err := p.client.Call(method, args...)

	if err != nil {
		log.Fatal(err)
	}

	return response
}
