package m

import (
	"reflect"
	"testing"
)

func TestNoteGroup_Dedupe(t *testing.T) {
	tests := []struct {
		name       string
		notes      NoteGroup
		wantResult NoteGroup
	}{
		{
			notes:      Group(3, 2, 1, 1, 2, 3),
			wantResult: Group(3, 2, 1),
		},
		{
			notes:      Group(3, 3, 3, 2, 2, 2, 1, 1, 1),
			wantResult: Group(3, 2, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := tt.notes.Dedupe(); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("NoteGroup.Dedupe() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestNoteGroup_Interleave(t *testing.T) {
	type args struct {
		others []NoteGroup
	}
	tests := []struct {
		name       string
		notes      NoteGroup
		args       args
		wantResult NoteGroup
	}{
		{
			name:  "different lengths",
			notes: Group(A, A),
			args: args{
				others: []NoteGroup{
					Group(B, B),
					Group(C, D, E),
				},
			},
			wantResult: Group(A, B, C, A, B, D),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := tt.notes.Interleave(tt.args.others...); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("NoteGroup.Interleave() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestNoteGroup_Over(t *testing.T) {
	type args struct {
		root NoteNumber
	}
	tests := []struct {
		name       string
		notes      NoteGroup
		args       args
		wantResult NoteGroup
	}{
		{
			notes:      Group(4, 5, 6, 7, 5, 4, 5, 3),
			args:       args{root: 5},
			wantResult: Group(5, 6, 7, 5, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := tt.notes.Over(tt.args.root); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("NoteGroup.Over() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestNoteGroup_Subgroup(t *testing.T) {
	type args struct {
		index int
		size  int
	}
	tests := []struct {
		name       string
		notes      NoteGroup
		args       args
		wantResult NoteGroup
	}{
		{
			name:       "Positive start index",
			notes:      Group(0, 10, 20, 30, 40, 50, 60, 70, 80, 90),
			args:       args{index: 1, size: 3},
			wantResult: Group(10, 20, 30),
		},
		{
			name:       "Positive start index, exceed length",
			notes:      Group(0, 10, 20, 30, 40, 50, 60, 70, 80, 90),
			args:       args{index: 8, size: 3},
			wantResult: Group(80, 90),
		},
		{
			name:       "Negative start index",
			notes:      Group(0, 10, 20, 30, 40, 50, 60, 70, 80, 90),
			args:       args{index: -1, size: 3},
			wantResult: Group(70, 80, 90),
		},
		{
			name:       "Negative start index, exceed range",
			notes:      Group(0, 10, 20, 30),
			args:       args{index: -3, size: 3},
			wantResult: Group(0, 10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := tt.notes.Subgroup(tt.args.index, tt.args.size); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("NoteGroup.Subgroup() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestNoteGroup_AllSubgroups(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name  string
		notes NoteGroup
		args  args
		want  []NoteGroup
	}{
		{
			notes: Group(0, 1, 2, 3, 4),
			args:  args{3},
			want:  []NoteGroup{Group(0, 1, 2), Group(1, 2, 3), Group(2, 3, 4)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.notes.AllSubgroups(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoteGroup.AllSubgroups() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoteGroup_Append(t *testing.T) {
	type args struct {
		appendages []NoteGroup
	}
	tests := []struct {
		name       string
		notes      NoteGroup
		args       args
		wantResult NoteGroup
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := tt.notes.Append(tt.args.appendages...); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("NoteGroup.Append() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestAppend(t *testing.T) {
	type args struct {
		appendages []NoteGroup
	}
	tests := []struct {
		name       string
		args       args
		wantResult NoteGroup
	}{
		{
			args:       args{[]NoteGroup{Group(1, 2), Group(3, 4), Group(5)}},
			wantResult: Group(1, 2, 3, 4, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Append(tt.args.appendages...); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Append() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
