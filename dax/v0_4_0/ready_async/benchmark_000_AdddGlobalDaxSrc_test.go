package ready_async_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/v0_4_0/ready_async"
)

func Benchmark_AddGlobalDaxSrc_zeroDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.FreeGlobalDaxSrcForTest("cliargs")
	}
	b.StopTimer()
}

func Benchmark_AddGlobalDaxSrc_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddGlobalDaxSrc("cliargs", FooDaxSrc{})
		sabi.FreeGlobalDaxSrcForTest("cliargs")
	}
	b.StopTimer()
}

func Benchmark_AddGlobalDaxSrc_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddGlobalDaxSrc("cliargs", FooDaxSrc{})
		sabi.AddGlobalDaxSrc("database", FooDaxSrc{})
		sabi.AddGlobalDaxSrc("pubsub", FooDaxSrc{})
		sabi.AddGlobalDaxSrc("json", FooDaxSrc{})
		sabi.AddGlobalDaxSrc("env", FooDaxSrc{})

		sabi.FreeGlobalDaxSrcForTest("cliargs")
		sabi.FreeGlobalDaxSrcForTest("database")
		sabi.FreeGlobalDaxSrcForTest("pubsub")
		sabi.FreeGlobalDaxSrcForTest("json")
		sabi.FreeGlobalDaxSrcForTest("env")
	}
	b.StopTimer()
}
