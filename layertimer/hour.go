package layertimer

import (
	"time"
	"container/list"
)

const (
	hour_slots uint8 = 24
)

type HourScale struct {
	Pos           uint64          // 槽位置
	CallBack                      // 回调函数
	Slot          []*list.List    // 槽定义
	SlotNum       uint8           // 槽数量
	Accuracy      time.Duration   // 精度
	Ticker        *time.Ticker    // 计时器
	TickerStop    chan struct{}   // 停止定时器
	Next          *HourScale      // 指针，指向下一个HourScale
}

func NewHourScale(ticker time.Duration) *HourScale {
	var pos = uint64(0)
	var accuracy = ticker * time.Hour
	return &HourScale {
		Pos: pos,
		Slot: make([]*list.List, hour_slots),
		SlotNum: hour_slots,
		Accuracy: accuracy,
		Ticker: time.NewTicker(accuracy),
		TickerStop: make(chan struct{}),
	}
}
