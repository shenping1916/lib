package layertimer

import (
	"fmt"
	"time"
	"sync/atomic"
	"container/list"
)

const (
	week_slots uint16 = 7
)

type CallBack = func(args ...interface{})

type WeekScale struct {
	Pos           uint64          // 槽位置
	CallBack                      // 回调函数
	Slot          []*list.List    // 槽定义
	Accuracy      time.Duration   // 精度
	Ticker        *time.Ticker    // 计时器
	TickerStop    chan struct{}   // 停止定时器
	Next          *WeekScale      // 指针，指向下一个DailyScale
}

func NewWeekScale(ticker time.Duration) *WeekScale {
	var pos = uint64(0)
	var accuracy = ticker * time.Hour * 24
	return &WeekScale {
		Pos: pos,
		Slot: make([]*list.List, week_slots),
		Accuracy: accuracy,
		Ticker: time.NewTicker(accuracy),
		TickerStop: make(chan struct{}),
	}
}

func (t *WeekScale) GetCurrentPos() uint64 {
	//var current_pos = t.Pos
	//return ((t. << 24) | (t.Pos & 0xffffff))
	return t.Pos
}


func (t *WeekScale) Start() {
	for {
		select {
		case <- t.Ticker.C:
			t.Handler()
		case <- t.TickerStop:
			t.Ticker.Stop()
			return
		}
	}
}

func (t *WeekScale) Stop() {
	t.TickerStop <- struct {}{}
}

func (t *WeekScale) Handler() {
	atomic.AddUint64(&t.Pos, 1)
	fmt.Println(time.Now().String(), t.Pos)
}



