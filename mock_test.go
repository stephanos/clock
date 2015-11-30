package clock

import (
	. "github.com/101loops/bdd"
	"time"
)

var (
	fixedTime = time.Unix(1415926535, 0)
)

var _ = Describe("Mock Clock", func() {

	It("returns the time", func() {
		clock := NewMock()
		Check(clock.IsFrozen(), IsFalse)
		Check(timeDiff(clock), IsLessThan, 1*time.Millisecond)
	})

	It("sets time", func() {
		clock := NewMock().Set(fixedTime)
		Check(clock.Now().Sub(fixedTime), IsLessThan, 1*time.Millisecond)

		time.Sleep(delay)
		Check(clock.Now().Sub(fixedTime), IsRoughly, delay, threshold)
	})

	It("adds time when not frozen", func() {
		clock := NewMock().Add(1 * time.Hour)
		Check(timeDiff(clock), IsRoughly, -1*time.Hour, threshold)
	})

	It("adds time when frozen", func() {
		clock := NewMock().Freeze()
		clock.Add(1 * time.Hour)
		Check(timeDiff(clock), IsRoughly, -1*time.Hour, threshold)
	})

	It("freezes", func() {
		clock := NewMock().Add(1 * time.Hour).Freeze()
		Check(clock.IsFrozen(), IsTrue)
		clockNow := clock.Now()

		time.Sleep(delay)

		Check(clock.Now(), IsSameTimeAs, clockNow)
	})

	It("freezes at passed-in time", func() {
		clock := NewMock().FreezeAt(fixedTime)
		Check(clock.IsFrozen(), IsTrue)

		time.Sleep(delay)

		Check(clock.Now(), IsSameTimeAs, fixedTime)
	})

	It("unfreezes", func() {
		clock := NewMock().Add(1 * time.Hour).Freeze()
		Check(clock.IsFrozen(), IsTrue)

		time.Sleep(delay)

		clock.Unfreeze()
		Check(clock.IsFrozen(), IsFalse)
		Check(timeDiff(clock), IsRoughly, -1*time.Hour+delay, threshold)
	})

	It("can sleep", func() {
		clock := NewMock()

		slept := durationOf(func() { clock.Sleep(delay) })
		Check(slept, IsRoughly, delay, threshold)
	})

	It("disables sleep", func() {
		clock := NewMock().NoSleep()

		slept := durationOf(func() { clock.Sleep(delay) })
		Check(slept, IsLessThan, delay)
	})

	It("overwrites sleep argument", func() {
		clock := NewMock().SetSleep(delay)

		slept := durationOf(func() { clock.Sleep(2 * time.Second) })
		Check(slept, IsRoughly, delay, threshold)
	})

	It("resets sleep override", func() {
		clock := NewMock().SetSleep(2 * time.Second)
		clock.ResetSleep()

		slept := durationOf(func() { clock.Sleep(delay) })
		Check(slept, IsRoughly, delay, threshold)
	})
})
