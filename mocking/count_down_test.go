package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountDown(t *testing.T) {
	t.Run("sequence test", func(t *testing.T) {
		spySleepPrinter := &SpyCountDownOperations{}
		// @param io.Writer: spySleepPrinter
		// @param Sleeper: spySleepPrinter
		CountDown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})

	t.Run("sleep test", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		CountDown(buffer, &SpyCountDownOperations{})
		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
