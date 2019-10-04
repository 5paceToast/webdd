package middleware

import (
	"github.com/valyala/fasthttp"
)

var (
	_ = Middleware(HTTPLog) // ensure compliance
)

// HTTPLog logs HTTP requests
func HTTPLog(ctx *fasthttp.RequestCtx) bool {
	ctx.Logger().Printf("[%q] %q", ctx.Method(), ctx.Path())
	return true
}
