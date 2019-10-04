package middleware

import (
	"os"
	"path"
	"reflect"
	"strings"
	"testing"

	"github.com/valyala/fasthttp"
	"toast.cafe/x/webdd/util"
)

func TestMarkdownFile(t *testing.T) {
	var ctx fasthttp.RequestCtx
	fpath := path.Join("testdata", "test.md")
	file, err := os.Open(fpath)
	if err != nil {
		t.Errorf("Failed to open file %q", fpath)
	}
	markdownFile(file, &ctx)

	ct := ctx.Response.Header.ContentType()
	if string(ct) != "text/html" {
		t.Errorf("Unexpected content type: %q", ct)
	}

	bd := string(ctx.Response.Body())
	if !strings.Contains(bd, "<p>Test data.</p>") {
		t.Errorf("Could not find test pargraph in output")
	}
}

func TestMarkDown(t *testing.T) {
	var md, html, wrong fasthttp.RequestCtx
	olddir := util.S.Dir
	util.S.Dir, _ = os.Getwd()
	defer func() { util.S.Dir = olddir }()

	md.Request.SetRequestURI("/testdata/test.md")
	html.Request.SetRequestURI("/testdata/test.html")
	wrong.Request.SetRequestURI("/testdata/test")

	mdres := MarkDown(&md)
	htmlres := MarkDown(&html)
	wrongres := MarkDown(&wrong)

	if mdres {
		t.Errorf("MarkDown claims it wants to continue parsing after path %q", md.Path())
	}
	if htmlres {
		t.Errorf("MarkDown claims it wants to continue parsing after path %q", html.Path())
	}
	if !wrongres {
		t.Errorf("MarkDown claims it successfully parsed path %q", wrong.Path())
	}

	if !reflect.DeepEqual(md.Response.Body(), html.Response.Body()) {
		t.Errorf("Parsing %q gives a different result from %q", md.Path(), html.Path())
	}
}

var benchmarkMarkDownResult bool

func benchmarkMarkDown(ctx *fasthttp.RequestCtx, b *testing.B) {
	var res bool
	olddir := util.S.Dir
	util.S.Dir, _ = os.Getwd()
	defer func() { util.S.Dir = olddir }()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res = MarkDown(ctx)
	}
	b.StopTimer()
	benchmarkMarkDownResult = res
}

func BenchmarkMarkDownFile(b *testing.B) {
	var ctx fasthttp.RequestCtx
	ctx.Request.SetRequestURI("/testdata/test.html")
	benchmarkMarkDown(&ctx, b)
}

func BenchmarkMarkDownPass(b *testing.B) {
	var ctx fasthttp.RequestCtx
	ctx.Request.SetRequestURI("/testdata/nonexistent.html")
	benchmarkMarkDown(&ctx, b)
}

func BenchmarkMarkDownSkip(b *testing.B) {
	var ctx fasthttp.RequestCtx
	ctx.Request.SetRequestURI("/testdata/skip")
	benchmarkMarkDown(&ctx, b)
}
