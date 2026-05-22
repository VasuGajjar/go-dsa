package medium

import "testing"

var subarraySumTests = []struct {
	name string
	nums []int
	k    int
	want int
}{
	{"example", []int{1, 1, 1}, 2, 2},
	{"single match", []int{3}, 3, 1},
	{"single no match", []int{3}, 2, 0},
	{"whole array", []int{1, 2, 3}, 6, 1},
	{"two subarrays", []int{1, 2, 3}, 3, 2},
	{"no valid subarray", []int{1, 2, 3}, 7, 0},
	{"with negatives", []int{1, -1, 1}, 1, 3},
	{"k zero with zeros", []int{0, 0, 0}, 0, 6},
	{"k zero with cancel pairs", []int{1, -1, 2, -2}, 0, 3},
	{"negative k", []int{1, -2, 3}, -2, 1},
	{"all same", []int{2, 2, 2, 2}, 4, 3},
	{"large k no match", []int{1, 2, 3}, 100, 0},
}

func TestSubarraySum(t *testing.T) {
	for _, tt := range subarraySumTests {
		t.Run(tt.name, func(t *testing.T) {
			got := SubarraySum(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("SubarraySum(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func BenchmarkSubarraySum(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for b.Loop() {
		SubarraySum(nums, 10)
	}
}
