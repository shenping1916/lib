package layertimer

import (
	"time"
	"testing"
)

func TestTask_NeedRound(t *testing.T) {
	task := Task{DelayTime: time.Duration(2374) * time.Second}
	wheel := SecondTimeWheel{}
	t.Log(wheel.RoundConSume())
	t.Log(task.NeedRound(wheel.RoundConSume()))
}

func TestTask_DesSlotsIndex(t *testing.T) {
	task := Task{DelayTime: time.Duration(67) * time.Second}

	wheel := SecondTimeWheel{}
	t.Log(task.DesSlotsIndex(wheel.RoundConSume(), uint64(3)))
}