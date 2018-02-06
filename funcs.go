package m

import (
	"fmt"
	"strings"
)

// Group creates a new NoteGroup with the supplied notes
func Group(notes ...NoteNumber) NoteGroup {
	return notes
}

// MinorTriad creates a minor triad based on root
func MinorTriad(root NoteNumber) NoteGroup {
	return Group(root, root+3, root+7)
}

// MajorTriad creates a minor triad based on root.
func MajorTriad(root NoteNumber) NoteGroup {
	return Group(root, root+4, root+7)
}

// MajorChord returns a major chord based on root including the upper octave
func MajorChord(root NoteNumber) NoteGroup {
	return MajorTriad(root).Append(Group(root + 12))
}

// MinorChord returns a minor chord based on root including the upper octave
func MinorChord(root NoteNumber) NoteGroup {
	return MinorTriad(root).Append(Group(root + 12))
}

// Sus4Triad creates a sus4 triad based on root
func Sus4Triad(root NoteNumber) NoteGroup {
	return Group(root, root+5, root+7)
}

// MajorSeventh chord based on root. Four notes long
func MajorSeventh(root NoteNumber) NoteGroup {
	return Group(root, root+4, root+7, root+11)
}

// MinorSeventh chord based on root. Four notes long
func MinorSeventh(root NoteNumber) NoteGroup {
	return Group(root, root+3, root+7, root+10)
}

// Dedupe iterates over a NoteGroup, and removes any notes that are not occuring
// for the first time.
func (notes NoteGroup) Dedupe() (result NoteGroup) {
	noteMap := make(map[NoteNumber]int, len(notes))
	result = make(NoteGroup, 0, len(notes))

	for _, note := range notes {
		noteMap[note]++
		if noteMap[note] == 1 {
			result = append(result, note)
		}
	}
	return result
}

// AllOctaves creates a new NoteGroup, with all the original notes repeated in
// every octave.
func (notes NoteGroup) AllOctaves() NoteGroup {
	result := make(NoteGroup, 0, 128)
	var i NoteNumber
	for i = lowestNote; i <= highestNote; i++ {
		for _, n := range notes.Dedupe() {
			if i%12 == n%12 {
				result = append(result, i)
				break
			}
		}
	}
	return result
}

// Append NoteGroups into one larger group
func (notes NoteGroup) Append(appendages ...NoteGroup) (result NoteGroup) {
	result = notes
	for _, group := range appendages {
		result = append(result, group...)
	}
	return result
}

// Interleave multiple groups together. This chooses the shortest group, and
// creates a new group with all of the others interleaved.
//
// In this example, there are 3 groups, and the shortest group has a length of
// 2, so the result is 6 units long:
//
// ([1,1]).Interleave([2,2], [5,6,7,8]) == [1,2,5,1,2,6]
func (notes NoteGroup) Interleave(others ...NoteGroup) (result NoteGroup) {
	// find the shortest group
	shortest := notes
	for _, group := range others {
		if len(group) < len(shortest) {
			shortest = group
		}
	}
	totalGroups := len(others) + 1
	groupSize := len(shortest)
	resultSize := groupSize * totalGroups
	result = make(NoteGroup, 0, resultSize)

	for i := 0; i < groupSize; i++ {
		result = append(result, notes[i])
		for _, group := range others {
			result = append(result, group[i])
		}
	}

	return result
}

// Over removes all notes below root. It does not remove root.
func (notes NoteGroup) Over(root NoteNumber) (result NoteGroup) {
	result = make(NoteGroup, 0, len(notes))
	for _, note := range notes {
		if note >= root {
			result = append(result, note)
		}
	}
	return result
}

// Under removes all notes above topNote. It does not remove topNote.
func (notes NoteGroup) Under(topNote NoteNumber) (result NoteGroup) {
	result = make(NoteGroup, 0, len(notes))
	for _, note := range notes {
		if note <= topNote {
			result = append(result, note)
		}
	}
	return result
}

// Subgroup reslices a note group given a size and starting index. Negative
// indices may be used to take a subgroup from the end of the slice.
// For example index=-1, size=3 give the last three elements from the group.
func (notes NoteGroup) Subgroup(index int, size int) (result NoteGroup) {
	if index >= 0 {
		end := index + size
		if end > len(notes) { // make sure we don't go over
			end = len(notes)
		}
		return notes[index:end]
	}

	// We have a negative index
	if size >= len(notes) {
		return notes
	}

	start := len(notes) - size + index + 1
	end := start + size
	if start < 0 {
		start = 0
	}
	if end > len(notes) {
		end = len(notes)
	}
	if end < 0 {
		end = 0
	}

	return notes[start:end]
}

// AllSubgroups return every subgroup with the specified size. For example:
//[1, 2, 3, 4].AllSubGroups(2) == [[1,2],[2,3],[3,4]]
func (notes NoteGroup) AllSubgroups(size int) []NoteGroup {
	if size < 0 {
		panic("AllSubgroups size must be greater than 0")
	}
	count := len(notes) - size + 1
	if count < 1 {
		return []NoteGroup{Group()}
	}
	result := make([]NoteGroup, count)
	for i := range result {
		result[i] = notes.Subgroup(i, size)
	}
	return result
}

// Append joins together multple NoteGroups. It is similar to NoteGroup.Append
func Append(appendages ...NoteGroup) (result NoteGroup) {
	size := 0
	for _, group := range appendages {
		size += len(group)
	}

	result = make(NoteGroup, 0, size)

	for _, group := range appendages {
		result = append(result, group...)
	}
	return result
}

// FirstOfEach returns a new NoteGroup with the first Note of each group
func FirstOfEach(groups ...NoteGroup) (result NoteGroup) {
	result = make(NoteGroup, len(groups))
	for i, group := range groups {
		result[i] = group[0]
	}
	return
}

// FlatName gets the scientific pitch name for a note. It chooses the flat name
// for notes not in the C-major scale.
func FlatName(n NoteNumber) string {
	root := n % 12
	name := pitchesFlats[root]
	octave := (int(n) / 12) - 1
	return fmt.Sprintf("%s%d", name, octave)
}

// FlatString make a string with all note names
func (notes NoteGroup) FlatString() (result string) {
	names := make([]string, len(notes))
	for i, n := range notes {
		names[i] = FlatName(n)
	}
	return "[" + strings.Join(names, " ") + "]"
}

// Repeat the NoteGroup `count` times.
func (notes NoteGroup) Repeat(count int) (result NoteGroup) {
	initialLength := len(notes)
	result = make(NoteGroup, initialLength*count)
	for i := range result {
		result[i] = notes[i%initialLength]
	}
	return result
}

// Transpose a NoteGroup up or down.
func (notes NoteGroup) Transpose(amt int) (result NoteGroup) {
	result = make(NoteGroup, len(notes))
	for i, note := range notes {
		result[i] = note + uint8(amt)
	}
	return
}

// Reverse the note group
func (notes NoteGroup) Reverse() (result NoteGroup) {
	result = make(NoteGroup, len(notes))
	for i, note := range notes {
		result[len(notes)-i-1] = note
	}
	return result
}
