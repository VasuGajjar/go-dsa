package medium

import (
	"fmt"
	"math/rand"
	"slices"
	"testing"
)

var longestConsecutiveTests = []struct {
	name string
	nums []int
	want int
}{
	{"empty", []int{}, 0},
	{"single element", []int{1}, 1},
	{"all same", []int{3, 3, 3}, 1},
	{"already sorted consecutive", []int{1, 2, 3, 4}, 4},
	{"reverse sorted consecutive", []int{4, 3, 2, 1}, 4},
	{"example leetcode", []int{100, 4, 200, 1, 3, 2}, 4},
	{"with duplicates", []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
	{"no consecutive", []int{1, 3, 5, 7}, 1},
	{"negative range", []int{-3, -2, -1, 0, 1}, 5},
	{"two disjoint sequences", []int{1, 2, 3, 10, 11, 12, 13}, 4},
	{"duplicates mixed", []int{1, 2, 2, 3, 3, 4}, 4},
	{"negative and positive", []int{-1, 0, 1, 2}, 4},
}

func TestLongestConsecutiveWithMap(t *testing.T) {
	for _, tt := range longestConsecutiveTests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestConsecutiveWithMap(tt.nums)
			if got != tt.want {
				t.Errorf("LongestConsecutiveWithMap(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestLongestConsecutiveWithSort(t *testing.T) {
	for _, tt := range longestConsecutiveTests {
		t.Run(tt.name, func(t *testing.T) {
			nums := slices.Clone(tt.nums)
			got := LongestConsecutiveWithSort(nums)
			if got != tt.want {
				t.Errorf("LongestConsecutiveWithSort(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkLongestConsecutiveWithMap(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	for _, n := range sizes {
		data := rand.Perm(n)
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for b.Loop() {
				LongestConsecutiveWithMap(data)
			}
		})
	}
}

func BenchmarkLongestConsecutiveWithSort(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	for _, n := range sizes {
		data := rand.Perm(n)
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for b.Loop() {
				input := make([]int, n)
				copy(input, data)
				LongestConsecutiveWithSort(input)
			}
		})
	}
}
