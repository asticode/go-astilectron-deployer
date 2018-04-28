package astideployer

import (
	"context"
	"time"

	"github.com/asticode/go-astitools/http"
	"github.com/julienschmidt/httprouter"
)

// ServePrivate serves the private routes
func (v *Deployer) ServePrivate(ctx context.Context) error {
	// Create router
	r := httprouter.New()

	// TODO Web

	// Chain middlewares
	h := astihttp.ChainMiddlewares(r,
		astihttp.MiddlewareBasicAuth(v.c.ServerPrivate.Username, v.c.ServerPrivate.Password),
		astihttp.MiddlewareTimeout(5*time.Second),
	)

	// Serve
	return astihttp.Serve(ctx, h, v.c.ServerPrivate.Addr, nil)
}
