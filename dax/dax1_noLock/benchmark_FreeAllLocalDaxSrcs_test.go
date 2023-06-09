package dax1_noLock_test

import (
	sabi "github.com/sttk-go/benchmarks_sabi/dax/dax1_noLock"
	"testing"
)

func Benchmark_FreeAllLocalDaxSrcs_oneDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddLocalDaxSrcForTest(base, "foo", FooDaxSrc{})
		base.FreeAllLocalDaxSrcs()
	}
	b.StopTimer()
}

func Benchmark_FreeAllLocalDaxSrcs_fiveDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddLocalDaxSrcForTest(base, "foo", FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "bar", BarDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "baz", BazDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "qux", QuxDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "corge", CorgeDaxSrc{})

		base.FreeAllLocalDaxSrcs()
	}
	b.StopTimer()
}
