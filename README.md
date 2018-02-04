# M

MIDI Centric Music Theory Library

Provides the NoteNumber and NoteGroup types.

To Create a c minor arpeggio from C2 to C6:

Midi note numbers correspond to [Scientific Pitch Notation](https://en.wikipedia.org/wiki/Scientific_pitch_notation).

The library provides at least one go `const` value for every midi note.

- Midi note 0 (aka C -1) is identified by `C`
- Midi note 12 (aka C 0) is identified by `C0`
- Midi note 13 (aka C# 0, aka Bb 0) is identified by `Cs0`, `Bf0`, and `Bb0`

```Go
// Create [0, 3, 7]
cMinor := m.MinorTriad(m.C)

// Create [36 39 43 48 51 55 60 63 67 72 75 79 84]
arp := cMinor.AllOctaves().Over(m.C2).Under(m.C6)
```
