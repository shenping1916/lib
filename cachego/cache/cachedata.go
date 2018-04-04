package cache

import (
	"time"
	"github.com/golang/glog"
)

const (
	// Never expire
	// When key expiration time NoExpire,
	// the key can not and should not be deleted, unless the flush cachetable
	NoExpire time.Duration = -1

	// default expiration time
	// when the key is not set expiration time,
	// the default expiration time setting
	// default: 30m
	DefaultExpire time.Duration = 30
)

type SetExpiration func(*Expiration)

type Expiration struct {
	Expire          time.Duration
}

type CacheData struct {
	Key             interface{}
	Value           interface{}
	Create          time.Duration
	Expire          time.Duration
}

func WithExpiration(expire time.Duration) SetExpiration {
	return func(o *Expiration) {
		o.Expire = expire
	}
}

func NewCacheData(key, value interface{}, opts ...SetExpiration) *CacheData {
	now := time.Now()
	options := Expiration{
		Expire: time.Duration(DefaultExpire) * time.Minute,
	}
	for _, o := range opts {
		o(&options)
	}

	return &CacheData {
		Key: key,
		Value: value,
		Create: time.Duration(now.Nanosecond()),
		Expire: time.Duration(now.Add(options.Expire).Nanosecond()),
	}
}

func (t *CacheTable) Add(key, value interface{}, expire time.Duration) {
	if !t.NotFound(key) {
		t.CacheData = NewCacheData(key, value, WithExpiration(expire))
		t.M.Store(key, t.CacheData)


	}
}

func (t *CacheTable) Update(key, value interface{}, expire time.Duration) {
	if t.NotFound(key) {
		glog.Error(ErrorKey)
		return
	}

	if !t.Expired(key) {
		t.Delete(key)
	}

	t.CacheData = NewCacheData(key, value, WithExpiration(expire))
	t.M.Store(key, t.CacheData)
}

func (t *CacheTable) Delete(keys... interface{}) {
	done := make(chan struct{})
	go func() {
		for _, key := range keys {
			t.M.Delete(key)
		}
		done <- struct{}{}
	}()

	<-done
}

func (t *CacheTable) Expired(key interface{}) bool {
	value := t.Value(key)
	if value != nil {
		if v, ok := value.(*CacheData); ok {
			if v.Expire == -1 {
				return true
			}

			return v.Expire.Nanoseconds() > time.Now().UnixNano()
		}
	}
}

func (t *CacheTable) NotFound(key interface{}) bool {
	_, ok := t.M.Load(key)
	return ok
}

func (t *CacheTable) Value(key interface{}) interface{} {
	value, ok := t.M.Load(key)
	if ok {
		return value
	}

	return nil
}

func (t *CacheTable) TableFlush(key interface{}) int {
	size := 0
	for _, table := range Ring.CacheTables {
		size += table.TableSize()
	}

	Ring.CacheTables = make([]*CacheTable, 0, maxtable)
	return size
}

func (t *CacheTable) TableSize() int {
	length := 0
	t.M.Range(func(_, _ interface{}) bool {
		length ++
		return true
	})

	return length
}

func (t *CacheTable) TriggerAddCallBack(f *func(*CacheData)) {
	select {
	case t.Addcallback <- t.CacheData:

	}
}