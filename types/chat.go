package types

import "time"

type Chat struct {
	CreatedTime   time.Time
	UserA         User
	UserB         User
	MessageStream []Message
}
