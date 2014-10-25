package clock

import . "github.com/101loops/bdd"

var _ = Describe("Clock Work", func() {

	It("Now()", func() {
		Work = New()
		Check(Now().Sub(now()), IsLessThan, delay)
	})

	It("Sleep()", func() {
		slept := durationOf(func() { Sleep(delay) })
		Check(slept, IsRoughly, delay, threshold)
	})

	It("After()", func() {
		elapsed := durationOf(func() { <-After(delay) })
		Check(elapsed, IsRoughly, delay, threshold)
	})

	It("Ticker()", func() {
		ticker := Ticker(delay).C
		elapsed := durationOf(func() { <-ticker; <-ticker })
		Check(elapsed, IsRoughly, 2*delay, threshold)
	})

	It("Tick()", func() {
		elapsed := durationOf(func() { <-Tick(delay); <-Tick(delay) })
		Check(elapsed, IsRoughly, 2*delay, threshold)
	})
})
