package easy

import "testing"

func TestMaxConsecutiveOne(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"ones and zeros mixed", []int{1, 1, 0, 1, 1, 1}, 3},
		{"single streak", []int{1, 0, 1, 1, 0, 1}, 2},
		{"all ones", []int{1, 1, 1, 1}, 4},
		{"all zeros", []int{0, 0, 0}, 0},
		{"single one", []int{1}, 1},
		{"single zero", []int{0}, 0},
		{"ones at start", []int{1, 1, 1, 0, 0}, 3},
		{"ones at end", []int{0, 0, 1, 1, 1}, 3},
		{"alternating", []int{1, 0, 1, 0, 1}, 1},
		{"empty", []int{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxConsecutiveOne(tt.input)
			if got != tt.want {
				t.Errorf("MaxConsecutiveOne(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
