package easy

import "testing"

func TestIsSorted(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  bool
	}{
		{"sorted ascending", []int{1, 2, 3, 4, 5}, true},
		{"sorted equal", []int{3, 3, 3}, true},
		{"unsorted", []int{3, 1, 2}, false},
		{"reverse sorted", []int{5, 4, 3, 2, 1}, false},
		{"single element", []int{42}, true},
		{"empty", []int{}, true},
		{"negatives sorted", []int{-5, -3, -1, 0, 2}, true},
		{"negatives unsorted", []int{-1, -5, 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsSorted(tt.input)
			if got != tt.want {
				t.Errorf("IsSorted(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
