package middleware

import (
	"bytes"
	"log"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

func testHTTPLog(buf *bytes.Buffer, ctx *fasthttp.RequestCtx, logger fasthttp.Logger, t *testing.T) {
	var expect = "[\"GET\"] \"/test\""
	ctx.Request.SetRequestURI("/test")
	HTTPLog(ctx)
	if !strings.Contains(buf.String(), expect) {
		t.Errorf("expected %q in buffer, got %q", expect, buf)
	}
}

func TestHTTPLog(t *testing.T) {
	var (
		buf    bytes.Buffer
		ctx    fasthttp.RequestCtx
		logger = log.New(&buf, "", log.Lshortfile)
	)
	ctx.Init(&ctx.Request, nil, logger)
	testHTTPLog(&buf, &ctx, logger, t)
}

var benchmarkHTTPLogResult bool

func benchmarkHTTPLog(ctx *fasthttp.RequestCtx, logger fasthttp.Logger, b *testing.B) {
	var res bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res = HTTPLog(ctx)
	}
	b.StopTimer()
	benchmarkHTTPLogResult = res
}

func BenchmarkHTTPLog(b *testing.B) {
	var (
		buf    bytes.Buffer
		ctx    fasthttp.RequestCtx
		logger = log.New(&buf, "", log.Lshortfile)
	)
	ctx.Init(&ctx.Request, nil, logger)
	ctx.Request.SetRequestURI("/test")
	benchmarkHTTPLog(&ctx, logger, b)
}
func BenchmarkHTTPLogrus(b *testing.B) {
	var (
		buf    bytes.Buffer
		ctx    fasthttp.RequestCtx
		logger = logrus.New()
	)
	logger.Out = &buf
	ctx.Init(&ctx.Request, nil, logger)
	ctx.Request.SetRequestURI("/test")
	benchmarkHTTPLog(&ctx, logger, b)
}
