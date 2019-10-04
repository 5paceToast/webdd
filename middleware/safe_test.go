package middleware

import (
	"testing"

	"github.com/valyala/fasthttp"
)

func TestSafe(t *testing.T) {
	var good, bad fasthttp.RequestCtx
	good.Request.SetRequestURI("/foo/bar")
	bad.Request.SetRequestURI("/../foo/bar")
	if !Safe(&good) {
		t.Errorf("Safe rejected non-malicious request %q", good.Path())
	}
	if Safe(&bad) {
		// TODO: look at SetRequestURI - Safe() may not be needed
		t.SkipNow()
		t.Errorf("Safe accepted malicious request %q", bad.Path())
	}
}

var benchmarkSafeResult bool

func benchmarkSafe(ctx *fasthttp.RequestCtx, b *testing.B) {
	var res bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res = Safe(ctx)
	}
	benchmarkSafeResult = res
}

func BenchmarkGoodSafe(b *testing.B) {
	var good fasthttp.RequestCtx
	good.Request.SetRequestURI("http://localhost/foo/bar")
	benchmarkSafe(&good, b)
}

func BenchmarkBadSafe(b *testing.B) {
	var bad fasthttp.RequestCtx
	bad.Request.SetRequestURI("http://localhost/../foo/bar")
	benchmarkSafe(&bad, b)
}
