package provider

import (
	"net/rpc/jsonrpc"
	"net/rpc"
	
	"log"
	"golang.org/x/net/websocket"
)

type Provider struct {
	config *Config
	client *rpc.Client
}


func New(c *Config) (p *Provider, err error) {
	p = new(Provider)
	p.config = c

	err = p.init(c)
	
	return p, err
}

func (p *Provider) init(c *Config) error {
	ws, err := websocket.Dial(c.URL, "", c.Origin)

	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	p.client = jsonrpc.NewClient(ws)

	log.Println("WS Connection Established")

	return err;
}

func (p *Provider) Close() error {
	return p.client.Close()
}