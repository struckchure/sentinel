package sentinel

import "fmt"

type IGateway interface {
	Run() error
}

type Gateway struct {
	config Config
}

func (g *Gateway) Run() error {
	fmt.Println(g.config)
	fmt.Println("running ran, but no implementation")

	return nil
}

func NewGateway(config Config) IGateway {
	return &Gateway{config: config}
}
