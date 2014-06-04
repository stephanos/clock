package clock

import (
	. "github.com/101loops/bdd"
	"testing"
	"time"
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

func durationOf(fn func()) time.Duration {
	beforeSleep := now()
	fn()
	return now().Sub(beforeSleep)
}
