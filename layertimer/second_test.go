package layertimer

import (
	"testing"
)

func TestNewSecondScale(t *testing.T) {
	second := NewSecondTimeWheel()
	go second.Start()
}
