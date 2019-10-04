package util

import (
	"testing"
)

func TestReplaceSuffix(t *testing.T) {
	var (
		s      = []byte("prefix.suffix")
		before = []byte(".suffix")
		after  = []byte(".different")
		expect = []byte("prefix.different")
	)
	res := ReplaceSuffix(s, before, after)
	if string(res) != string(expect) {
		t.Errorf("Expected %q from ReplaceSuffix but got %q", expect, res)
	}
}

var benchmarkReplaceSuffixResult []byte

func BenchmarkReplaceSuffix(b *testing.B) {
	var (
		s      = []byte("prefix.suffix")
		before = []byte(".suffix")
		after  = []byte(".different")
		res    []byte
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res = ReplaceSuffix(s, before, after)
	}
	b.StopTimer()
	benchmarkReplaceSuffixResult = res
}
