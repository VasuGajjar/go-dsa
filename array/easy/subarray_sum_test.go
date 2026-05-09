package easy

import "testing"

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"longest not at start", []int{1, 2, 3, 4, 5}, 9, 3},            // [2,3,4]=9, len 3
		{"entire array", []int{1, 2, 3, 4}, 10, 4},                      // whole array sums to 10
		{"starts at zero with negatives", []int{1, -1, 5, -2, 3}, 3, 4}, // [1,-1,5,-2]=3, len 4
		{"subarray not at start", []int{2, 1, 2}, 3, 2},                 // [2,1] or [1,2], len 2
		{"single match", []int{3}, 3, 1},
		{"single no match", []int{1}, 2, 0},
		{"no match", []int{1, 2, 3}, 10, 0},
		{"empty array", []int{}, 5, 0},
		{"negative numbers k zero", []int{-1, -2, 1, 2}, 0, 4}, // whole array sums to 0
		{"k zero with zeros", []int{0, 0, 0}, 0, 3},
		{"all same elements", []int{1, 1, 1, 1}, 2, 2},
		{"negative k", []int{-3, -2, -1, 4, 5}, -3, 2}, // [-2,-1]=-3, len 2
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SubarraySum(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("SubarraySum(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
