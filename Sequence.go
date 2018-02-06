package m

import (
	"fmt"
	"sort"
	"time"

	"github.com/CharlesHolbrow/gm"
)

// An Event is just a interface{} type. You can add any time to a sequence
type Event interface{}

// SequenceEvent wraps an event, and adds metadata about the location.
type SequenceEvent struct {
	subPosition int
	position    float64       // Dimensionless floating point value relative to start
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
	content map[float64]int

	// list is where the actual sorting happens
	list []SequenceEvent

	// Sequences have a cursor, which saves a point within the sequence.
	// Sequence methods may use the point to inform behavior. For example, the
	// .Get method uses the curor as a loop point.
	Cursor float64

	// Is sequence.list known to be in playback order?
	sorted bool
}

// NewSequence creates and initializes a new Sequence
func NewSequence() *Sequence {
	return &Sequence{
		list:    make([]SequenceEvent, 0),
		content: make(map[float64]int),
	}
}

// Add an event to the sequence. Position is a dimensionless point to place the
// event. The dimension can be set with the sequence.Sorted() function.
func (s *Sequence) Add(position float64, event Event) *Sequence {
	if position < 0 {
		fmt.Printf("Bad event position: %f (%v)\n", position, event)
		panic("Cannot add event to with negative position")
	}

	s.sorted = false

	timeEvent := SequenceEvent{
		Event:       event,
		position:    position + s.Cursor,
		subPosition: s.content[position],
	}

	s.content[position]++
	s.list = append(s.list, timeEvent)

	return s
}

// AddSubdivisions creates `n` sustain events euqally spaced over `totalLength`
// between the start of the sequence and the cursor.
func (s *Sequence) AddSubdivisions(n int, totalLength, duty float64) *Sequence {
	if totalLength == 0 {
		panic("Sequence.AddSubdivisons requires non-aero `totalLength`")
	}

	spacing := totalLength / float64(n)
	length := spacing * duty
	for i := 0; i < n; i++ {
		s.AddSustain(float64(i)*spacing, length, 100)
	}
	return s
}

// RampSustainVelocity replaces velocity value in sustained events. The new
// velocity values ramp from startVel at pos=0 to endVal at pos=s.Cursor
func (s *Sequence) RampSustainVelocity(startVel, endVel int) {
	slope := float64(endVel-startVel) / s.Cursor
	for i, sEvent := range s.list {
		if e, ok := sEvent.Event.(gm.Note); ok {
			position := sEvent.position
			e.Vel = uint8(slope*position + float64(startVel))
			s.list[i].Event = e
		}
	}
}

// AddSustain adds an event with a Non-zero length.
func (s *Sequence) AddSustain(position, length float64, velocity int) *Sequence {
	// For now, I'm using gm.Note events for sustained events. It might be
	// advantagous to use something more speciffic so that this doesn't get
	// confused for a midi sequence. If I decide to change the event type
	// make sure to also update the `AddRhythmicMelody` implementation.
	s.Add(position, gm.Note{Vel: uint8(velocity), On: true})
	s.list[len(s.list)-1].length = length
	return s
}

// AddRhythmicMelody add a melody with a given rhythm.
// To use:
// - Create a rhythmic sequence `r`
// - Call `r.AddSustain(...)` one or more times
// - Set the loop point by setting `r.Cursor`
// - Create a NoteGroup with the desired Melody
// - On the receiver sequence, `s`, call `s.AddRythmicMelody(...)`
func (s *Sequence) AddRhythmicMelody(rhythm *Sequence, notes NoteGroup, midiCh int) *Sequence {
	ch := uint8(midiCh)
	for i, root := range notes {
		seqEvent := rhythm.Get(i)
		length := seqEvent.Length()
		if note, ok := seqEvent.Event.(gm.Note); ok && length > 0 {
			onPos := seqEvent.Position()
			offPos := onPos + length
			s.Add(onPos, gm.Note{On: true, Note: root, Ch: ch, Vel: note.Vel})
			s.Add(offPos, gm.Note{Note: root, Ch: ch})
		}
	}
	return s
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

// Get an event by its index, looping from the beginning of the sequence to the
// cursor. If the are events after the cursor of the sequence, the order of
// events returned by Get may not follow the playback order. To get the playback
// order, use EventList method.
func (s *Sequence) Get(i int) SequenceEvent {
	if !s.sorted {
		s.sort()
	}

	if s.Cursor == 0 {
		fmt.Println("Sequence.Get - WARNING - zero cursor")
	}

	repetition := i / len(s.list)
	event := s.list[i%len(s.list)]
	event.position = event.position + float64(repetition)*s.Cursor
	return event
}

// Play back the sequence on the supplied channel. If out is nil, create a
// channel. returns the playback channel.
func (s *Sequence) Play(unit time.Duration) chan interface{} {
	start := time.Now()
	s.sort()
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
