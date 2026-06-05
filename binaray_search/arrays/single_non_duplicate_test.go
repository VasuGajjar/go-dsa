package arrays

import "testing"

func TestSingleNonDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"single element", []int{1}, 1},
		{"unique at start", []int{1, 2, 2, 3, 3}, 1},
		{"unique at end", []int{1, 1, 2, 2, 3}, 3},
		{"unique in middle", []int{1, 1, 2, 3, 3}, 2},
		{"unique at mid odd index", []int{1, 1, 2, 2, 3, 3, 4}, 4},
		{"unique at first mid", []int{1, 2, 2, 3, 3, 4, 4}, 1},
		{"large odd position", []int{0, 0, 1, 1, 2, 3, 3, 4, 4}, 2},
		{"large even position", []int{0, 1, 1, 2, 2, 3, 3, 4, 4}, 0},
		{"negative numbers", []int{-3, -3, -1, -1, 0}, 0},
		{"two pairs unique last", []int{1, 1, 2}, 2},
		{"two pairs unique first", []int{1, 2, 2}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SingleNonDuplicate(tt.nums)
			if got != tt.want {
				t.Errorf("SingleNonDuplicate(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
