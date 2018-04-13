package layertimer

import (
	"time"
	"container/list"
)

const (
	minute_slots uint8 = 60
)

type MinuteScale struct {
	Pos           uint64          // 槽位置
	CallBack                      // 回调函数
	Slot          []*list.List    // 槽定义
	SlotNum       uint8           // 槽数量
	Accuracy      time.Duration   // 精度
	Ticker        *time.Ticker    // 计时器
	TickerStop    chan struct{}   // 停止定时器
	Next          *MinuteScale    // 指针，指向下一个MinuteScale
}

func NewMinuteScale(ticker time.Duration) *MinuteScale {
	var pos = uint64(0)
	var accuracy = ticker * time.Minute
	return &MinuteScale {
		Pos: pos,
		Slot: make([]*list.List, minute_slots),
		SlotNum: minute_slots,
		Accuracy: accuracy,
		Ticker: time.NewTicker(accuracy),
		TickerStop: make(chan struct{}),
	}
}
