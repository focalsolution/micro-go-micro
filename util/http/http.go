package http

import (
	"net/http"

	"github.com/focalsolution/micro-go-micro/client/selector"
	"github.com/focalsolution/micro-go-micro/registry"
)

func NewRoundTripper(opts ...Option) http.RoundTripper {
	options := Options{
		Registry: registry.DefaultRegistry,
	}
	for _, o := range opts {
		o(&options)
	}

	return &roundTripper{
		rt:   http.DefaultTransport,
		st:   selector.Random,
		opts: options,
	}
}
