package medium

import (
	"slices"
	"testing"
)

var nextPermutationTests = []struct {
	name string
	nums []int
	want []int
}{
	{"example 1", []int{1, 2, 3}, []int{1, 3, 2}},
	{"example 2", []int{3, 2, 1}, []int{1, 2, 3}},
	{"example 3", []int{1, 1, 5}, []int{1, 5, 1}},
	{"single element", []int{1}, []int{1}},
	{"two elements ascending", []int{1, 2}, []int{2, 1}},
	{"two elements descending", []int{2, 1}, []int{1, 2}},
	{"already largest", []int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
	{"already smallest", []int{1, 2, 3, 4}, []int{1, 2, 4, 3}},
	{"duplicates", []int{1, 3, 1, 3}, []int{1, 3, 3, 1}},
	{"all same", []int{2, 2, 2}, []int{2, 2, 2}},
	{"mid permutation", []int{2, 3, 1}, []int{3, 1, 2}},
	{"longer sequence", []int{1, 3, 5, 4, 2}, []int{1, 4, 2, 3, 5}},
}

func TestNextPermutation(t *testing.T) {
	for _, tt := range nextPermutationTests {
		t.Run(tt.name, func(t *testing.T) {
			nums := slices.Clone(tt.nums)
			NextPermutation(nums)
			if !slices.Equal(nums, tt.want) {
				t.Errorf("NextPermutation(%v) = %v, want %v", tt.nums, nums, tt.want)
			}
		})
	}
}
