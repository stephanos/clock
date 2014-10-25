package clock

import . "github.com/101loops/bdd"

var _ = Describe("Real Clock", func() {

	var clock = New()

	It("Now()", func() {
		clockNow := clock.Now()
		Check(clockNow.Sub(now()), IsLessThan, delay)
	})

	It("Sleep()", func() {
		slept := durationOf(func() { clock.Sleep(delay) })
		Check(slept, IsRoughly, delay, threshold)
	})

	It("After()", func() {
		elapsed := durationOf(func() { <-clock.After(delay) })
		Check(elapsed, IsRoughly, delay, threshold)
	})

	It("Ticker()", func() {
		ticker := clock.Ticker(delay).C
		elapsed := durationOf(func() { <-ticker; <-ticker })
		Check(elapsed, IsRoughly, 2*delay, threshold)
	})

	It("Tick()", func() {
		elapsed := durationOf(func() { <-clock.Tick(delay); <-clock.Tick(delay) })
		Check(elapsed, IsRoughly, 2*delay, threshold)
	})
})
