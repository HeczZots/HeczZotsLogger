package logger

import "math/rand"

type Sampler interface {
	// Sample returns true if the event should be part of the sample, false if
	// the event should be dropped.
	Sample(lvl Level) bool
}
type RandomSampler uint32

// Sample implements the Sampler interface.
func (s RandomSampler) Sample(lvl Level) bool {
	if s <= 0 {
		return false
	}
	if rand.Intn(int(s)) != 0 {
		return false
	}
	return true
}
