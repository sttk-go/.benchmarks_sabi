package ready_sync_test

import (
	sabi "github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0/ready_sync"
	"sync"
)

type FooDaxConn struct{}

func (conn FooDaxConn) Commit(wg sync.WaitGroup) sabi.Err { return sabi.Ok() }
func (conn FooDaxConn) Committed() sabi.Err               { return sabi.Ok() }
func (conn FooDaxConn) Rollback(wg sync.WaitGroup)        {}
func (conn FooDaxConn) Close()                            {}

type FooDaxSrc struct{}

func (ds FooDaxSrc) SetUp(wg sync.WaitGroup) sabi.Err { return sabi.Ok() }
func (ds FooDaxSrc) Ready() sabi.Err                  { return sabi.Ok() }
func (ds FooDaxSrc) End()                             {}
func (ds FooDaxSrc) CreateDaxConn() (sabi.DaxConn, sabi.Err) {
	return FooDaxConn{}, sabi.Ok()
}
