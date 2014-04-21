package clock

import "time"

type realClock struct{}

// Real is a Clock that wraps the behaviour of the time package.
// It can not be manipulated.
var Real Clock

func init() {
	Real = &realClock{}
}

func (*realClock) Now() time.Time {
	return time.Now()
}

func (*realClock) Sleep(d time.Duration) {
	time.Sleep(d)
}

func (*realClock) Tick(d time.Duration) <-chan time.Time {
	return time.Tick(d)
}

func (*realClock) Ticker(d time.Duration) *time.Ticker {
	return time.NewTicker(d)
}

func (*realClock) After(d time.Duration) <-chan time.Time {
	return time.After(d)
}

