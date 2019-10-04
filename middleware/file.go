package middleware

import (
	"github.com/valyala/fasthttp"
	"toast.cafe/x/webdd/util"
)

var (
	f *fasthttp.FS
	h fasthttp.RequestHandler
	_ = Middleware(File) // ensure compliance
)

// File serves the file if it exists
func File(ctx *fasthttp.RequestCtx) bool {
	if f == nil {
		f = &fasthttp.FS{
			Root:               util.S.Dir,
			IndexNames:         []string{"index.html"},
			GenerateIndexPages: true,
			Compress:           false,
			AcceptByteRange:    true,
		}
		h = f.NewRequestHandler()
	}
	h(ctx)
	return false
}
