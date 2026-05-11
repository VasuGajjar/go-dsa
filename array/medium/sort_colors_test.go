package medium

import (
	"slices"
	"testing"
)

var sortColorsTests = []struct {
	name string
	nums []int
	want []int
}{
	{"all zeros", []int{0, 0, 0}, []int{0, 0, 0}},
	{"all ones", []int{1, 1, 1}, []int{1, 1, 1}},
	{"all twos", []int{2, 2, 2}, []int{2, 2, 2}},
	{"already sorted", []int{0, 1, 2}, []int{0, 1, 2}},
	{"reverse sorted", []int{2, 1, 0}, []int{0, 1, 2}},
	{"mixed", []int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
	{"single element", []int{1}, []int{1}},
	{"two elements", []int{2, 0}, []int{0, 2}},
	{"no ones", []int{2, 0, 2, 0}, []int{0, 0, 2, 2}},
	{"no zeros", []int{2, 1, 2, 1}, []int{1, 1, 2, 2}},
	{"no twos", []int{1, 0, 1, 0}, []int{0, 0, 1, 1}},
}

func TestSortColorsWithCount(t *testing.T) {
	for _, tt := range sortColorsTests {
		t.Run(tt.name, func(t *testing.T) {
			nums := slices.Clone(tt.nums)
			SortColorsWithCount(nums)
			if !slices.Equal(nums, tt.want) {
				t.Errorf("SortColorsWithCount(%v) = %v, want %v", tt.nums, nums, tt.want)
			}
		})
	}
}

func TestSortColorsWithDutchFlagAlgorithm(t *testing.T) {
	for _, tt := range sortColorsTests {
		t.Run(tt.name, func(t *testing.T) {
			nums := slices.Clone(tt.nums)
			SortColorsWithDutchFlagAlgorithm(nums)
			if !slices.Equal(nums, tt.want) {
				t.Errorf("SortColorsWithDutchFlagAlgorithm(%v) = %v, want %v", tt.nums, nums, tt.want)
			}
		})
	}
}

func BenchmarkSortColorsWithCount(b *testing.B) {
	nums := []int{2, 0, 2, 1, 1, 0, 2, 1, 0, 2, 0, 1}
	for b.Loop() {
		input := slices.Clone(nums)
		SortColorsWithCount(input)
	}
}

func BenchmarkSortColorsWithDutchFlagAlgorithm(b *testing.B) {
	nums := []int{2, 0, 2, 1, 1, 0, 2, 1, 0, 2, 0, 1}
	for b.Loop() {
		input := slices.Clone(nums)
		SortColorsWithDutchFlagAlgorithm(input)
	}
}
