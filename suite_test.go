package clock

import (
	"testing"
	"time"
	. "github.com/101loops/bdd"
)

var (
	delay     = 200 * time.Millisecond
	threshold = 100 * time.Millisecond
)

func TestSuite(t *testing.T) {
	RunSpecs(t, "Clock Suite")
}

func now() time.Time {
	return time.Now()
}

func timeDiff(c Clock) time.Duration {
	return time.Now().Sub(c.Now())
}

func durationOf(fn func()) time.Duration {
	beforeSleep := now()
	fn()
	return now().Sub(beforeSleep)
}
