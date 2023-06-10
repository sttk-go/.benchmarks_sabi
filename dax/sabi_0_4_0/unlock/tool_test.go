package unlock_test

import (
	sabi "github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0/unlock"
)

type FooDaxConn struct{}

func (conn FooDaxConn) Commit() sabi.Err { return sabi.Ok() }
func (conn FooDaxConn) Rollback()        {}
func (conn FooDaxConn) Close()           {}

type FooDaxSrc struct{}

func (ds FooDaxSrc) SetUp() sabi.Err { return sabi.Ok() }
func (ds FooDaxSrc) End()            {}
func (ds FooDaxSrc) CreateDaxConn() (sabi.DaxConn, sabi.Err) {
	return FooDaxConn{}, sabi.Ok()
}
