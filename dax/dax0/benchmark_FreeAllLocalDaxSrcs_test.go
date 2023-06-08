package dax0_test

import (
	sabi "github.com/sttk-go/benchmarks_sabi/dax/dax0"
	"testing"
)

func Benchmark_FreeAllLocalDaxSrcs_oneDs(b *testing.B) {
  base := sabi.NewDaxBase()
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    base.SetUpLocalDaxSrc("foo", FooDaxSrc{})
    base.FreeAllLocalDaxSrcs()
  }
}

func Benchmark_FreeAllLocalDaxSrcs_fiveDs(b *testing.B) {
  base := sabi.NewDaxBase()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    base.SetUpLocalDaxSrc("foo", FooDaxSrc{})
    base.SetUpLocalDaxSrc("bar", BarDaxSrc{})
    base.SetUpLocalDaxSrc("baz", BazDaxSrc{})
    base.SetUpLocalDaxSrc("qux", QuxDaxSrc{})
    base.SetUpLocalDaxSrc("corge", CorgeDaxSrc{})

    base.FreeAllLocalDaxSrcs()
  }
}
