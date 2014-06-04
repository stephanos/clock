package clock

import . "github.com/101loops/bdd"

var _ = Describe("Clock Work", func() {

	Context("Now()", func() {
		Work = New()
		Check(Now().Sub(now()), IsLessThan, delay)
	})

	Context("Sleep()", func() {
		slept := durationOf(func() { Sleep(delay) })
		Check(slept, IsRoughly, delay, threshold)
	})

	Context("After()", func() {
		elapsed := durationOf(func() { <-After(delay) })
		Check(elapsed, IsRoughly, delay, threshold)
	})

	Context("Ticker()", func() {
		ticker := Ticker(delay).C
		elapsed := durationOf(func() { <-ticker; <-ticker })
		Check(elapsed, IsRoughly, 2*delay, threshold)
	})

	Context("Tick()", func() {
		elapsed := durationOf(func() { <-Tick(delay); <-Tick(delay) })
		Check(elapsed, IsRoughly, 2*delay, threshold)
	})
})
