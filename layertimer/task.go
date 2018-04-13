package layertimer

import (
	"time"
)

type callback = func(args ...interface{})

type SetTask = func(*TaskOption)

type TaskOption struct {
	DelayTime     time.Duration
	Once          bool
}

type Task struct {
	UUID          string          // 任务uuid
	DelayTime     time.Duration   // 任务延迟执行时间
	Once          bool            // 是否只执行一次，默认: false(执行后删除)，否则不删除
	Circle        int             // 时间轮需转动圈数
	CallBack      callback        // 回调函数
}

func WithTaskSet(delay time.Duration, once bool) SetTask {
	return func(o *TaskOption) {
		o.DelayTime = delay
		o.Once = once
	}
}

func NewTask(delay time.Duration, once bool, cb CallBack, roundcousume time.Duration, opts ...SetTask) *Task{
	uuid := GenertorUUID()
	options := TaskOption{
		DelayTime: delay,
		Once: once,
	}

	for _, o := range opts {
		o(&options)
	}

	task := &Task {
		UUID: uuid,
		DelayTime: options.DelayTime,
		Once: options.Once,
		CallBack: cb,
	}
	task.Circle = task.NeedRound(roundcousume)

	return task
}

func (t *Task) NeedRound(roundcousume time.Duration) int {
	delay := t.DelayTime
	return int(delay / roundcousume)
}

func (t *Task) DesSlotsIndex(roundcousume time.Duration, current_pos uint64) []uint64 {
	slots_ := make([]uint64, 0, SlotsNum)
	delay := t.DelayTime
	for delay < roundcousume && roundcousume % delay == 0 {
		index := uint64(roundcousume / delay)
		slots_ = append(slots_, index)
		return slots_
	}

	index_ := current_pos + uint64(delay % roundcousume)
	slots_ = append(slots_, index_)
	return slots_
}
