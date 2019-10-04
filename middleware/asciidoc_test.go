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

func TestAsciidocFile(t *testing.T) {
	var ctx fasthttp.RequestCtx
	fpath := path.Join("testdata", "test.adoc")
	file, err := os.Open(fpath)
	if err != nil {
		t.Errorf("Failed to open file %q", fpath)
	}
	asciidocFile(file, &ctx)

	ct := ctx.Response.Header.ContentType()
	if string(ct) != "text/html" {
		t.Errorf("Unexpected content type: %q", ct)
	}

	bd := string(ctx.Response.Body())
	if !strings.Contains(bd, "<p>Test data.</p>") {
		t.Errorf("Could not find test pargraph in output")
	}
}

func TestASCIIDoc(t *testing.T) {
	var adoc, html, wrong fasthttp.RequestCtx
	olddir := util.S.Dir
	util.S.Dir, _ = os.Getwd()
	defer func() { util.S.Dir = olddir }()

	adoc.Request.SetRequestURI("/testdata/test.adoc")
	html.Request.SetRequestURI("/testdata/test.html")
	wrong.Request.SetRequestURI("/testdata/test")

	adocres := ASCIIDoc(&adoc)
	htmlres := ASCIIDoc(&html)
	wrongres := ASCIIDoc(&wrong)

	if adocres {
		t.Errorf("ASCIIDoc claims it wants to continue parsing after path %q", adoc.Path())
	}
	if htmlres {
		t.Errorf("ASCIIDoc claims it wants to continue parsing after path %q", html.Path())
	}
	if !wrongres {
		t.Errorf("ASCIIDoc claims it successfully parsed path %q", wrong.Path())
	}

	if !reflect.DeepEqual(adoc.Response.Body(), html.Response.Body()) {
		t.Errorf("Parsing %q gives a different result from %q", adoc.Path(), html.Path())
	}
}

var benchmarkASCIIDocResult bool

func benchmarkASCIIDoc(ctx *fasthttp.RequestCtx, b *testing.B) {
	var res bool
	olddir := util.S.Dir
	util.S.Dir, _ = os.Getwd()
	defer func() { util.S.Dir = olddir }()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res = ASCIIDoc(ctx)
	}
	b.StopTimer()
	benchmarkASCIIDocResult = res
}

func BenchmarkASCIIDocFile(b *testing.B) {
	var ctx fasthttp.RequestCtx
	ctx.Request.SetRequestURI("/testdata/test.html")
	benchmarkASCIIDoc(&ctx, b)
}

func BenchmarkASCIIDocPass(b *testing.B) {
	var ctx fasthttp.RequestCtx
	ctx.Request.SetRequestURI("/testdata/nonexistent.html")
	benchmarkASCIIDoc(&ctx, b)
}

func BenchmarkASCIIDocSkip(b *testing.B) {
	var ctx fasthttp.RequestCtx
	ctx.Request.SetRequestURI("/testdata/skip")
	benchmarkASCIIDoc(&ctx, b)
}
