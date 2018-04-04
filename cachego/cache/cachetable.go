package cache

import (
	"sync"
	"sync/atomic"
)

var number uint32 = 0

type CacheTable struct {
	M             *sync.Map
	CacheData     *CacheData
	Addcallback   chan *CacheData
}

func UnAvailNumber() bool {
	n := atomic.LoadUint32(&number)
	return n > Ring.MaxTable
}

func NewCacheTable() *CacheTable {
	var m sync.Map
	if !UnAvailNumber() {
		table := &CacheTable{
			M: &m,
		}
		Ring.CacheTables = append(Ring.CacheTables, table)
		atomic.AddUint32(&number, 1)

		return table
	}

	return nil
}