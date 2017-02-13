// Clock stub file

// To use the right term, this is the package *clause*.
// You can document general stuff about the package here if you like.
package clock

import "fmt"

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

// Clock API as stub definitions.  No, it doesn't compile yet.
// More details and hints are in clock_test.go.

type Clock struct {
	minutes int
}

func New(hour, minute int) Clock { // -25
	// one day has 1440 minutes, and the day part can be omitted
	minutes := (hour*60 + minute) % 1440
	if minutes < 0 {
		minutes += 1440
	}

	return Clock{minutes}
}

func (c Clock) String() string {
	hour := c.minutes / 60 % 24
	minute := c.minutes % 60
	return fmt.Sprintf("%02d:%02d", hour, minute)

}

func (c Clock) Add(minutes int) Clock {
	m := c.minutes + minutes
	return New(0, m)

}
