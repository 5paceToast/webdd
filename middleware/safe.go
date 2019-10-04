package middleware

import (
	"bytes"
	"github.com/valyala/fasthttp"
)

var (
	_ = Middleware(Safe) // ensure compliance
)

// Safe refuses to continue if the path is deemed unsafe
func Safe(ctx *fasthttp.RequestCtx) bool {
	p := ctx.Path()
	if bytes.Contains(p, []byte{'.', '.'}) {
		return false // no escaping the directory
	}
	return true
}
