package middleware

import "github.com/valyala/fasthttp"

// Middleware is a function that takes a RequestCtx and returns a Bool (continue or not)
type Middleware func(ctx *fasthttp.RequestCtx) bool

// HandleMiddlewares returns a pure fasthttp handler that simply goes through every middleware in order
func HandleMiddlewares(ms []Middleware) func(ctx *fasthttp.RequestCtx) {
	f := func(ctx *fasthttp.RequestCtx) {
		for _, m := range ms {
			if !m(ctx) {
				break
			}
		}
	}
	return f
}
