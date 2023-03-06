package model

import (
	"time"
)

type Logger struct {
	ID             string
	Username       string    `bson:"username"`
	LoginTimestamp time.Time `bson:"login_timestamp"`
}
