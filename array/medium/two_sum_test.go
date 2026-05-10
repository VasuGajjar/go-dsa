package medium

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"pair at start", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"pair at end", []int{3, 2, 4}, 6, []int{1, 2}},
		{"same value twice", []int{3, 3}, 6, []int{0, 1}},
		{"pair not adjacent", []int{1, 5, 3, 2}, 4, []int{0, 2}},
		{"negative numbers", []int{-3, 4, 3, 90}, 0, []int{0, 2}},
		{"target with negatives", []int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
		{"no solution", []int{1, 2, 3}, 100, []int{-1, -1}},
		{"single element", []int{5}, 10, []int{-1, -1}},
		{"zero target", []int{0, 4, 3, 0}, 0, []int{0, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			if got[0] != tt.want[0] || got[1] != tt.want[1] {
				t.Errorf("TwoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
