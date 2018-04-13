package layertimer

import (
	"github.com/jakehl/goid"
)

func GenertorUUID() string {
	uuid := goid.NewV4UUID()
	return uuid.String()
}
