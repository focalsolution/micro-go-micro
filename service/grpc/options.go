package grpc

import (
	"crypto/tls"

	"github.com/focalsolution/micro-go-micro"
	gc "github.com/focalsolution/micro-go-micro/client/grpc"
	gs "github.com/focalsolution/micro-go-micro/server/grpc"
)

// WithTLS sets the TLS config for the service
func WithTLS(t *tls.Config) micro.Option {
	return func(o *micro.Options) {
		o.Client.Init(
			gc.AuthTLS(t),
		)
		o.Server.Init(
			gs.AuthTLS(t),
		)
	}
}
