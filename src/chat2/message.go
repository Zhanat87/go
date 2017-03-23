package chat2

import (
	"time"
)

// message represents a single message
type message struct {
	Name    string
	Message string
	When    time.Time
}
