package middleware

import (
	"bufio"
	"bytes"
	"context"
	"io"
	ospath "path"
	"os"

	"github.com/bytesparadise/libasciidoc"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/valyala/fasthttp"
	"toast.cafe/x/webdd/util"
)

var (
	_ = Middleware(ASCIIDoc) // ensure compliance
)

func asciidocFile(reader io.Reader, ctx *fasthttp.RequestCtx) {
	libasciidoc.ConvertToHTML(context.TODO(), reader, ctx, renderer.IncludeHeaderFooter(true))
	ctx.SetContentType("text/html")
}

// ASCIIDoc is a Middleware handler to catch AsciiDoc files
func ASCIIDoc(ctx *fasthttp.RequestCtx) bool {
	var (
		adoc = []byte{'.', 'a', 'd', 'o', 'c'}
		html = []byte{'.', 'h', 't', 'm', 'l'}
	)
	path := ctx.Path()
	switch {
	case bytes.HasSuffix(path, html):
		path = util.ReplaceSuffix(path, html, adoc)
		fallthrough
	case bytes.HasSuffix(path, adoc):
		fpath := ospath.Join(util.S.Dir, string(path))
		file, err := os.Open(fpath)
		if err == nil {
			bufr := bufio.NewReader(file)
			asciidocFile(bufr, ctx)
			return false
		}
	}
	return true
}
