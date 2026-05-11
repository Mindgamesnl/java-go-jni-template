package main

import "time"

type Timings struct {
	events []TimingEvent
}

type TimingEvent struct {
	Name     string
	Duration time.Duration
}

type PollableEvent struct {
	Name  string
	start time.Time
	t     *Timings
}

func (t *Timings) Measure(name string, fn func()) {
	start := time.Now()
	fn()
	duration := time.Since(start)
	t.events = append(t.events, TimingEvent{Name: name, Duration: duration})
}

func (t *Timings) StartEvent(name string) *PollableEvent {
	return &PollableEvent{Name: name, start: time.Now(), t: t}
}

func (e *PollableEvent) End() {
	duration := time.Since(e.start)
	e.t.events = append(e.t.events, TimingEvent{Name: e.Name, Duration: duration})
}

func NewTimings() *Timings {
	return &Timings{}
}
