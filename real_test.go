package clock

import . "github.com/101loops/bdd"

var _ = Describe("Real Clock", func() {

	var clock = New()

	Context("Now()", func() {
		clockNow := clock.Now()
		Check(clockNow.Sub(now()), IsLessThan, delay)
	})

	Context("Sleep()", func() {
		slept := durationOf(func() { clock.Sleep(delay) })
		Check(slept, IsRoughly, delay, threshold)
	})

	Context("Ticker()", func() {
		ticker := clock.Ticker(delay).C
		elapsed := durationOf(func() { <-ticker; <-ticker })
		Check(elapsed, IsRoughly, 2*delay, threshold)
	})

	Context("After()", func() {
		elapsed := durationOf(func() { <-clock.After(delay) })
		Check(elapsed, IsRoughly, delay, threshold)
	})

	Context("Tick()", func() {
		elapsed := durationOf(func() { <-clock.Tick(delay); <-clock.Tick(delay) })
		Check(elapsed, IsRoughly, 2*delay, threshold)
	})
})
