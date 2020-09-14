package provider

import (
	"log"

	"github.com/ybbus/jsonrpc"
)

type Provider struct {
	config *Config
	client jsonrpc.RPCClient
}

/*type Request struct {
	Method string `json:"method"`
	Params [1]interface{} `json:"params"`
	Id uint256 `json:"id"`
}*/

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
