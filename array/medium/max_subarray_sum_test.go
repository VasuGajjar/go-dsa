package medium

import "testing"

func TestMaxSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"single positive", []int{5}, 5},
		{"single negative", []int{-3}, 0},
		{"all positive", []int{1, 2, 3, 4}, 10},
		{"all negative", []int{-1, -2, -3}, 0},
		{"mixed positive dominant", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"mixed negative dominant", []int{-5, 1, -4}, 1},
		{"subarray in middle", []int{-3, 4, 5, -10, 2}, 9},
		{"subarray at end", []int{1, -5, 3, 4}, 7},
		{"subarray at start", []int{4, 3, -5, 1}, 7},
		{"single zero", []int{0}, 0},
		{"zeros and positives", []int{0, 3, 0, 2}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxSubarraySum(tt.nums)
			if got != tt.want {
				t.Errorf("MaxSubarraySum(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
