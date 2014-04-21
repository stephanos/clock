package clock

import (
	. "github.com/101loops/bdd"
	"time"
)

var (
	threshold = 100 * time.Millisecond
	fixedTime = time.Unix(1415926535, 0)
)

var _ = Describe("Mock Clock", func() {

	Context("Now()", func() {

		It("behaves like time.Now by default", func() {
			clock := NewMock()
			Check(clock.IsFrozen(), IsFalse)

			clockNow := clock.Now()
			Check(clockNow.Sub(now()), IsLessThan, 1*time.Millisecond)
		})

		It("sets time", func() {
			clock := NewMock().Set(fixedTime)

			Check(clock.Now().Sub(fixedTime), IsLessThan, 1*time.Millisecond)

			time.Sleep(1 * time.Second)
			Check(clock.Now().Sub(fixedTime), IsRoughly, 1*time.Second, threshold)
		})

		It("adds time", func() {
			clock := NewMock().Add(1 * time.Hour)

			clockNow := clock.Now()
			Check(clockNow.Sub(now()), IsRoughly, 1*time.Hour, threshold)
		})

		It("freezes", func() {
			clock := NewMock().Freeze()
			Check(clock.IsFrozen(), IsTrue)

			time.Sleep(1 * time.Second)
			Check(now().Sub(clock.Now()), IsRoughly, 1*time.Second, threshold)
		})

		It("freezes at a passed-in time", func() {
			clock := NewMock().FreezeAt(fixedTime)
			Check(clock.IsFrozen(), IsTrue)

			time.Sleep(1 * time.Second)
			Check(clock.Now(), IsSameTimeAs, fixedTime)
		})

		It("unfreezes", func() {
			clock := NewMock().Freeze()
			time.Sleep(1 * time.Second)

			clock.Unfreeze()
			Check(clock.IsFrozen(), IsFalse)
			Check(now().Sub(clock.Now()), IsRoughly, 1*time.Second, threshold)
		})
	})

	Context("Sleep()", func() {

		It("behaves like time.Sleep by default", func() {
			clock := NewMock()

			slept := durationOf(func() { clock.Sleep(1 * time.Second) })
			Check(slept, IsRoughly, 1*time.Second, threshold)
		})

		It("disables sleep", func() {
			clock := NewMock().NoSleep()

			slept := durationOf(func() { clock.Sleep(1 * time.Second) })
			Check(slept, IsLessThan, 1*time.Second)
		})

		It("overwrites sleep argument", func() {
			clock := NewMock().SetSleep(1 * time.Second)

			slept := durationOf(func() { clock.Sleep(2 * time.Second) })
			Check(slept, IsRoughly, 1*time.Second, threshold)
		})

		It("resets sleep override", func() {
			clock := NewMock().SetSleep(2 * time.Second)
			clock.ResetSleep()

			slept := durationOf(func() { clock.Sleep(1 * time.Second) })
			Check(slept, IsRoughly, 1*time.Second, threshold)
		})
	})
})

func now() time.Time {
	return time.Now()
}

func durationOf(fn func()) time.Duration {
	beforeSleep := now()
	fn()
	return now().Sub(beforeSleep)
}
