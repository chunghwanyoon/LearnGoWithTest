package mocking

import "time"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountDownOperations struct {
	Calls []string
}

// Sleep SpyCountDownOperations implement Sleeper interface
func (s *SpyCountDownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// Write SpyCountDownOperations implement io.Writer interface
func (s *SpyCountDownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(duration time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

const sleep = "sleep"
const write = "write"
