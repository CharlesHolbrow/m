package m

// C  = 0
// Db = 1
// D  = 2
// Eb = 3
// E  = 4
// F  = 5
// Gb = 6
// G  = 7
// Ab = 8
// A  = 9
// Bb = 10
// B  = 11

// C4 = 60
// C3 = 48
// C2 = 36
// C1 = 24
// C0 = 12

// NoteNumber is a Midi Note Number
type NoteNumber = uint8

// NoteGroup is a slice of NoteNumbers
type NoteGroup []NoteNumber

const lowestNote = 0
const highestNote = 127

// C  Pitch Class
const C = 0

// Cs represents the C Sharp Pitch Class
const Cs = 1

// Df represents the D Flat Pitch Class
const Df = 1

// Db represents the Flat Pitch Class
const Db = 1

// D Pitch Class
const D = 2

// Ds represents the D Sharp Pitch Class
const Ds = 3

// Ef represents E Flat Pitch Class
const Ef = 3

// Eb represents the E Flat Pitch Class
const Eb = 3

// E  Pitch Class
const E = 4

// F  Pitch Class
const F = 5

// Fs represents the F Sharp Pitch Class
const Fs = 6

// Gf represents the G Flat Pitch Class
const Gf = 6

// Gb represents the G Flat Pitch Class
const Gb = 6

// G represetns the G Pitch Class
const G = 7

// Gs Represents the G Sharp Pitch Class
const Gs = 8

// Af represents the A Flat Pitch Class
const Af = 8

// Ab represents the A Flat Pitch Class
const Ab = 8

// A  Pitch Class
const A = 9

// As represents the A  Sharp Pitch Class
const As = 10

// Bf represents the B Flat Pitch Class
const Bf = 10

// Bb represents the Flat Pitch Class
const Bb = 10

// B  Pitch Class
const B = 11

var pitchesFlats = [...]string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb"}
var pitchesSharps = [...]string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#"}
var pitchesMap = map[string]int{
	"C":  C,
	"C#": Cs, "Cs": Cs, "Db": Db, "Df": Db,
	"D":  D,
	"D#": Ds, "Ds": Ds, "Eb": Eb, "Ef": Ef,
	"E":  E,
	"F":  F,
	"F#": Fs, "Fs": Fs, "Gb": Gb, "Gf": Gf,
	"G":  G,
	"G#": Gs, "Gs": Gs, "Ab": Ab, "Af": Af,
	"A":  A,
	"A#": As, "As": As, "Bb": Bb, "Bf": Bf,
	"B": B}

// The MIDI Note Names below were generated automatically by my music theory
// namegen package "github.com/CharlesHolbrow/mt-namegen"

const C0 = 12
const C1 = 24
const C2 = 36
const C3 = 48
const C4 = 60
const C5 = 72
const C6 = 84
const C7 = 96
const C8 = 108
const C9 = 120
const Cs0 = 13
const Cs1 = 25
const Cs2 = 37
const Cs3 = 49
const Cs4 = 61
const Cs5 = 73
const Cs6 = 85
const Cs7 = 97
const Cs8 = 109
const Cs9 = 121
const Db0 = 13
const Db1 = 25
const Db2 = 37
const Db3 = 49
const Db4 = 61
const Db5 = 73
const Db6 = 85
const Db7 = 97
const Db8 = 109
const Db9 = 121
const Df0 = 13
const Df1 = 25
const Df2 = 37
const Df3 = 49
const Df4 = 61
const Df5 = 73
const Df6 = 85
const Df7 = 97
const Df8 = 109
const Df9 = 121
const D0 = 14
const D1 = 26
const D2 = 38
const D3 = 50
const D4 = 62
const D5 = 74
const D6 = 86
const D7 = 98
const D8 = 110
const D9 = 122
const Ds0 = 15
const Ds1 = 27
const Ds2 = 39
const Ds3 = 51
const Ds4 = 63
const Ds5 = 75
const Ds6 = 87
const Ds7 = 99
const Ds8 = 111
const Ds9 = 123
const Eb0 = 15
const Eb1 = 27
const Eb2 = 39
const Eb3 = 51
const Eb4 = 63
const Eb5 = 75
const Eb6 = 87
const Eb7 = 99
const Eb8 = 111
const Eb9 = 123
const Ef0 = 15
const Ef1 = 27
const Ef2 = 39
const Ef3 = 51
const Ef4 = 63
const Ef5 = 75
const Ef6 = 87
const Ef7 = 99
const Ef8 = 111
const Ef9 = 123
const E0 = 16
const E1 = 28
const E2 = 40
const E3 = 52
const E4 = 64
const E5 = 76
const E6 = 88
const E7 = 100
const E8 = 112
const E9 = 124
const F0 = 17
const F1 = 29
const F2 = 41
const F3 = 53
const F4 = 65
const F5 = 77
const F6 = 89
const F7 = 101
const F8 = 113
const F9 = 125
const Fs0 = 18
const Fs1 = 30
const Fs2 = 42
const Fs3 = 54
const Fs4 = 66
const Fs5 = 78
const Fs6 = 90
const Fs7 = 102
const Fs8 = 114
const Fs9 = 126
const Gb0 = 18
const Gb1 = 30
const Gb2 = 42
const Gb3 = 54
const Gb4 = 66
const Gb5 = 78
const Gb6 = 90
const Gb7 = 102
const Gb8 = 114
const Gb9 = 126
const Gf0 = 18
const Gf1 = 30
const Gf2 = 42
const Gf3 = 54
const Gf4 = 66
const Gf5 = 78
const Gf6 = 90
const Gf7 = 102
const Gf8 = 114
const Gf9 = 126
const G0 = 19
const G1 = 31
const G2 = 43
const G3 = 55
const G4 = 67
const G5 = 79
const G6 = 91
const G7 = 103
const G8 = 115
const G9 = 127
const Gs0 = 20
const Gs1 = 32
const Gs2 = 44
const Gs3 = 56
const Gs4 = 68
const Gs5 = 80
const Gs6 = 92
const Gs7 = 104
const Gs8 = 116
const Ab0 = 20
const Ab1 = 32
const Ab2 = 44
const Ab3 = 56
const Ab4 = 68
const Ab5 = 80
const Ab6 = 92
const Ab7 = 104
const Ab8 = 116
const Af0 = 20
const Af1 = 32
const Af2 = 44
const Af3 = 56
const Af4 = 68
const Af5 = 80
const Af6 = 92
const Af7 = 104
const Af8 = 116
const A0 = 21
const A1 = 33
const A2 = 45
const A3 = 57
const A4 = 69
const A5 = 81
const A6 = 93
const A7 = 105
const A8 = 117
const As0 = 22
const As1 = 34
const As2 = 46
const As3 = 58
const As4 = 70
const As5 = 82
const As6 = 94
const As7 = 106
const As8 = 118
const Bb0 = 22
const Bb1 = 34
const Bb2 = 46
const Bb3 = 58
const Bb4 = 70
const Bb5 = 82
const Bb6 = 94
const Bb7 = 106
const Bb8 = 118
const Bf0 = 22
const Bf1 = 34
const Bf2 = 46
const Bf3 = 58
const Bf4 = 70
const Bf5 = 82
const Bf6 = 94
const Bf7 = 106
const Bf8 = 118
const B0 = 23
const B1 = 35
const B2 = 47
const B3 = 59
const B4 = 71
const B5 = 83
const B6 = 95
const B7 = 107
const B8 = 119
