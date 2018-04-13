package layertimer

import (
	"time"
	"sync/atomic"
	"container/list"
	"sync"
)

var second_intelval time.Duration = 1

const (
	// 槽的总数
	SlotsNum uint64 = 60
)

type TaskOp interface {
	TaskAdd(task ...*Task)
	TaskRemove(id ...string) bool
	TaskMotify(id ...string) bool
}

type SecondTimeWheel struct {
	Cursor          uint64           // 游标
	Slots           []*SlotUint      // 所有槽
	Ticker          *time.Ticker     // 定时器
	TickInterval    time.Duration    // 定时器精度(时间刻度)
	TickerStop      chan struct{}    // 停止定时器
}

type SlotUint struct {
	Index           uint64           // 槽索引(编号)
	LinkList        *list.List       // 双链表
	*sync.Mutex                      // 互斥锁，确保操作双链表线程安全
	*Task                            // 任务指针，存储在双链表(LinkList)内
	TaskOpertion    TaskOp           // 实现TaskOp接口
	//Next            *SlotUint        // 指针，指向下一个槽(单元)
	//Prev            *SlotUint        // 指针，指向前一个槽(单元)
}

func NewSecondTimeWheel() *SecondTimeWheel {
	time_wheel := &SecondTimeWheel{
		Slots: make([]*SlotUint, SlotsNum),
		TickInterval: second_intelval * time.Second,
		TickerStop: make(chan struct{}),
	}

	time_wheel.Init()
	return time_wheel
}

func (t *SecondTimeWheel) Init() {
	slots := t.Slots
	for i :=0; i < len(slots); i ++ {
		slots[i] = &SlotUint{
			Index: uint64(i),
			LinkList: list.New(),
			Task: nil,
		}
	}
}

func (t *SecondTimeWheel) Start() {
	t.Ticker = time.NewTicker(second_intelval * time.Second)
	go t.Control()
}

func (t *SecondTimeWheel) Control() {
	Loop:
	for {
		select {
		case <- t.Ticker.C:
			atomic.AddUint64(&t.Cursor, 1)
			t.Handler()
		case <- t.TickerStop:
			t.Ticker.Stop()
			break Loop
		}
	}

	return
}

func (t *SecondTimeWheel) Stop() {
	t.TickerStop <- struct {}{}
}

func (t *SecondTimeWheel) GetCurrentCursor() uint64 {
	cursor := t.Cursor
	return cursor % (SlotsNum / uint64(second_intelval))
}

func (t *SecondTimeWheel) RoundConSume() time.Duration {
	return second_intelval * time.Duration(SlotsNum)
}

func (t *SecondTimeWheel) Handler() {
	var current_cursor = t.GetCurrentCursor()

	slotuint := t.Slots[current_cursor]
	l := slotuint.LinkList

	slotuint.Mutex.Lock()
	defer slotuint.Mutex.Unlock()
	for e := l.Front(); e != nil; {
		task := e.Value.(*Task)
		if task.Circle > 0 {
			task.Circle--
			e = e.Next()
			continue
		}

		go task.CallBack()
		next := e.Next()
		if task.Once {
			l.Remove(e)
		}
		e = next
	}
}

func (t *SecondTimeWheel) TaskAdd(task ...*Task) {
	if task == nil {
		return
	}

	for i :=0; i <= len(task); i++ {

	}
}

func (t *SecondTimeWheel) TaskRemove(id ...string) bool {
	for i :=0; i <= len(id); i++ {

	}
	return true
}

func (t *SecondTimeWheel) TaskMotify(id ...string) bool {
	for i :=0; i <= len(id); i++ {

	}
	return true
}
