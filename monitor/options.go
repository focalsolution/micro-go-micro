package monitor

import (
	"github.com/focalsolution/micro-go-micro/client"
	"github.com/focalsolution/micro-go-micro/registry"
)

type Options struct {
	Client   client.Client
	Registry registry.Registry
}

type Option func(*Options)

func Client(c client.Client) Option {
	return func(o *Options) {
		o.Client = c
	}
}

func Registry(r registry.Registry) Option {
	return func(o *Options) {
		o.Registry = r
	}
}
