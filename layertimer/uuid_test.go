package layertimer

import "testing"

func TestGenertorUUID(t *testing.T) {
	id := GenertorUUID()
	t.Log(id)
}
