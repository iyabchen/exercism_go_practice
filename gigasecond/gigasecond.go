// Package clause.
package gigasecond

import (
	"math"
	"time"
)

// Constant declaration.
const testVersion = 4 // find the value in gigasecond_test.go

// API function.  It uses a type from the Go standard library.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(math.Pow(10, 9)) * time.Duration(time.Second))
}
