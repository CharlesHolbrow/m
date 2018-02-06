package m

import (
	"fmt"
	"sort"
	"time"
)

// An Event is just a interface{} type. You can add any time to a sequence
type Event interface{}

// SequenceEvent wraps an event, and adds metadata about the location.
type SequenceEvent struct {
	subPosition int
	position    float64       // Dimensionless floating point value
	length      float64       // Sustained events have a non-zero length
	Time        time.Duration // Duration from seq start to event time
	Event       Event
}

// Position returns the dimensionless position of the event
func (se SequenceEvent) Position() float64 { return se.position }

// Length returns the dimensionless length of the event
func (se SequenceEvent) Length() float64 { return se.length }

// Sequence is an ordered collection of Events.
type Sequence struct {
	// content is important, because it stores the order that events were added
	content map[float64][]SequenceEvent

	// list is where the actual sorting happens
	list []SequenceEvent

	// How long is the sequence. Note: there may be events after the end of the
	// sequence (However, this should probably be avoided, as dangling note off
	// events could prematurely silence a subsequent note)
	length float64

	// Is sequence.list known to be in playback order?
	sorted bool
}

// NewSequence creates and initializes a new Sequence
func NewSequence(length float64) *Sequence {
	return &Sequence{
		list:    make([]SequenceEvent, 0),
		content: make(map[float64][]SequenceEvent),
		length:  length,
	}
}

// Get an event with Looping. If the are events after the end of the sequence,
// The order of events returned by Get may not follow the playback order. To get
// the playback order, use EventList method.
func (s *Sequence) Get(i int) SequenceEvent {
	if !s.sorted {
		s.sort()
	}

	repetition := i / len(s.list)
	event := s.list[i%len(s.list)]
	event.position = event.position + float64(repetition)*s.length
	return event
}

// Add an event to the sequence. Position is a dimensionless point to place the
// event. The dimension can be set with the sequence.Sorted() function.
func (s *Sequence) Add(position float64, event Event) {
	if position < 0 {
		fmt.Printf("Bad event position: %f (%v)\n", position, event)
		panic("Cannot add event to with negative position")
	}

	s.sorted = false

	events, ok := s.content[position]
	if !ok {
		events = make([]SequenceEvent, 0, 10)
		s.content[position] = events
	}

	timeEvent := SequenceEvent{
		Event:       event,
		position:    position,
		subPosition: len(events),
	}

	s.content[position] = append(events, timeEvent)
	s.list = append(s.list, timeEvent)
}

func (s *Sequence) AddSustain(position, length float64, event Event) {
	s.Add(position, event)
	s.content[position][len(s.content[position])-1].length = length
	s.list[len(s.list)-1].length = length
}

// EventList creates a slice of TimeEvents. The .Time property of each event will
// be populated. To Add an event, you had to specify a dimensionless time
// position. Set that dimension now with the `unit` argument.
func (s *Sequence) EventList(unit time.Duration) []SequenceEvent {
	s.sort()
	result := make([]SequenceEvent, len(s.list))
	for i, tEvent := range s.list {
		tEvent.Time = time.Duration(tEvent.position * float64(unit))
		result[i] = tEvent
	}
	return result
}

// Play back the sequence on the supplied channel. If out is nil, create a
// channel. returns the playback channel.
func (s *Sequence) Play(unit time.Duration) chan interface{} {
	start := time.Now()
	out := make(chan interface{})
	go func() {
		for _, tEvent := range s.EventList(unit) {
			time.Sleep(time.Until(start.Add(tEvent.Time)))
			out <- tEvent.Event
		}
		close(out)
	}()
	return out
}

// sort the underlying list
func (s *Sequence) sort() {
	if s.sorted {
		return
	}
	sort.Sort(s)
	s.sorted = true
}

// sort.Interface methods
func (s *Sequence) Len() int      { return len(s.list) }
func (s *Sequence) Swap(i, j int) { s.list[i], s.list[j] = s.list[j], s.list[i] }
func (s *Sequence) Less(i, j int) bool {
	if s.list[i].position == s.list[j].position {
		return s.list[i].subPosition < s.list[j].subPosition
	}
	return s.list[i].position < s.list[j].position
}
