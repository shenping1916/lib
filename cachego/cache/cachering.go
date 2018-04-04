package cache

var maxtable uint32 = 1024
var size uint64 = 1 << 32 -1

type CacheRing struct {
	MaxTable     uint32
	Size         uint64
	CacheTables  []*CacheTable
}

var Ring = &CacheRing{
	MaxTable: maxtable,
	Size: size,
	CacheTables: make([]*CacheTable, 0, maxtable),
}