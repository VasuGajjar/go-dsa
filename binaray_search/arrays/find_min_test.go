package arrays

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"not rotated", []int{1, 3, 5, 7, 9}, 1},
		{"rotated once", []int{3, 4, 5, 1, 2}, 1},
		{"rotated to last", []int{2, 3, 4, 5, 1}, 1},
		{"rotated to second", []int{5, 1, 2, 3, 4}, 1},
		{"fully rotated", []int{1, 2, 3, 4, 5}, 1},
		{"two elements rotated", []int{2, 1}, 1},
		{"two elements sorted", []int{1, 2}, 1},
		{"single element", []int{5}, 5},
		{"large rotation", []int{4, 5, 6, 7, 0, 1, 2}, 0},
		{"negative numbers", []int{-3, -2, -1, -5, -4}, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMin(tt.nums)
			if got != tt.want {
				t.Errorf("FindMin(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
