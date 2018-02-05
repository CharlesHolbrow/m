package m

// A Pattern is a rhythmic loop. Each entry in a pattern is a PatternEvent,
// which supports events that include a duration, such as midi notes.
//
// Events must be added in order.
type Pattern struct {
	duration float64
	cursor   float64
	events   []PatternEvent
}

// PatternEvent saves start/stop time like a Musical event
type PatternEvent struct {
	StartPosition float64
	Duration      float64
	Value         int
}

// NewPattern creates and initializes a new Pattern
func NewPattern(duration float64) *Pattern {
	return &Pattern{
		duration: duration,
		events:   make([]PatternEvent, 0),
	}
}

// NewPatternSubdivisions creates a pattern of duration 1. That sequence is
// divided it into n equally space regions. duty is the fraction of each region
// that will be active.
func NewPatternSubdivisions(n int, duty float64) *Pattern {
	p := NewPattern(1)
	spacing := float64(1) / float64(n)
	duration := spacing * duty

	for i := 0; i < n; i++ {
		p.Push(duration, 100).Advance(spacing)
	}
	return p
}

// Push an Event to the pattern. Chainable.
func (p *Pattern) Push(duration float64, value int) *Pattern {
	p.events = append(p.events, PatternEvent{
		StartPosition: p.cursor,
		Duration:      duration,
		Value:         value,
	})
	return p
}

// Advance the cursor.
func (p *Pattern) Advance(amt float64) *Pattern {
	p.cursor = p.cursor + amt
	return p
}

// Get a value
func (p *Pattern) Get(i int) PatternEvent {
	repitition := i / len(p.events)

	event := p.events[i%len(p.events)]
	event.StartPosition = event.StartPosition + (float64(repitition) * p.duration)

	return event
}

// RampValue adjust all the values in the pattern. It will linearly adjust each
// value based on it's start position.
func (p *Pattern) RampValue(start, end int) {
	slope := float64(end-start) / p.duration
	for i := 0; i < len(p.events); i++ {
		startPos := p.events[i].StartPosition
		value := int(float64(start) + slope*startPos)
		if value < end {
			value = end
		}
		p.events[i].Value = value
	}
}
