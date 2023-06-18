package ready_async_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/sabi_0_4_0/ready_async"
)

func Benchmark_FreeLocalDaxSrc_oneDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddLocalDaxSrcForTest(base, "cliargs", FooDaxSrc{})
		base.FreeLocalDaxSrc("cliargs")
	}
	b.StopTimer()
}

func Benchmark_FreeLocalDaxSrc_fiveDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddLocalDaxSrcForTest(base, "cliargs", FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "database", FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "pubsub", FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "json", FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "env", FooDaxSrc{})

		base.FreeLocalDaxSrc("cliargs")
		base.FreeLocalDaxSrc("database")
		base.FreeLocalDaxSrc("pubsub")
		base.FreeLocalDaxSrc("json")
		base.FreeLocalDaxSrc("env")
	}
	b.StopTimer()
}
