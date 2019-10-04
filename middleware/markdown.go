package middleware

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	ospath "path"

	"github.com/valyala/fasthttp"
	"github.com/yuin/goldmark"
	"toast.cafe/x/webdd/util"
)

var (
	_ = Middleware(MarkDown) // ensure compliance
)

func markdownFile(reader io.Reader, ctx *fasthttp.RequestCtx) {
	data, _ := ioutil.ReadAll(reader)
	goldmark.Convert(data, ctx)
	ctx.SetContentType("text/html")
}

// MarkDown is a Middleware handler to catch MarkDown files
func MarkDown(ctx *fasthttp.RequestCtx) bool {
	var (
		md   = []byte{'.', 'm', 'd'}
		html = []byte{'.', 'h', 't', 'm', 'l'}
	)
	path := ctx.Path()
	switch {
	case bytes.HasSuffix(path, html):
		path = util.ReplaceSuffix(path, html, md)
		fallthrough
	case bytes.HasSuffix(path, md):
		fpath := ospath.Join(util.S.Dir, string(path))
		file, err := os.Open(fpath)
		if err == nil {
			bufr := bufio.NewReader(file)
			markdownFile(bufr, ctx)
			return false
		}
	}
	return true
}
