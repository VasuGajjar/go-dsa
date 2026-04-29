package easy

import (
	"math"
	"testing"
)

func TestLargestElement(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"ascending", []int{1, 2, 3, 4, 5}, 5},
		{"descending", []int{5, 4, 3, 2, 1}, 5},
		{"single", []int{7}, 7},
		{"all equal", []int{4, 4, 4}, 4},
		{"negatives", []int{-3, -1, -2}, -1},
		{"mixed", []int{-5, 0, 3, -1, 9, 2}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LargestElement(tt.input)
			if got != tt.want {
				t.Errorf("LargestElement(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestSecondLargestElement(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"ascending", []int{1, 2, 3, 4, 5}, 4},
		{"descending", []int{5, 4, 3, 2, 1}, 4},
		{"two elements", []int{1, 2}, 1},
		{"negatives", []int{-5, -1, -3}, -3},
		{"mixed", []int{3, 1, 4, 1, 5, 9, 2, 6}, 6},
		{"single", []int{42}, math.MinInt},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SecondLargestElement(tt.input)
			if got != tt.want {
				t.Errorf("SecondLargestElement(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
