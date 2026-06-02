package arrays

import "testing"

func TestSearchRotatedI(t *testing.T) {
	var searchRotatedTests = []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"not rotated found", []int{1, 3, 5, 7, 9}, 5, 2},
		{"not rotated not found", []int{1, 3, 5, 7, 9}, 4, -1},
		{"rotated found left half", []int{4, 5, 6, 7, 0, 1, 2}, 5, 1},
		{"rotated found right half", []int{4, 5, 6, 7, 0, 1, 2}, 1, 5},
		{"rotated found pivot", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{"rotated not found", []int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{"rotated found first", []int{4, 5, 6, 7, 0, 1, 2}, 4, 0},
		{"rotated found last", []int{4, 5, 6, 7, 0, 1, 2}, 2, 6},
		{"single element match", []int{1}, 1, 0},
		{"single element no match", []int{1}, 2, -1},
		{"two elements rotated found", []int{3, 1}, 1, 1},
		{"two elements rotated not found", []int{3, 1}, 2, -1},
	}
	for _, tt := range searchRotatedTests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchRotatedI(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("SearchRotated(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func TestSearchRotatedII(t *testing.T) {
	var searchRotatedIITests = []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{"not rotated found", []int{1, 3, 5, 7, 9}, 5, true},
		{"not rotated not found", []int{1, 3, 5, 7, 9}, 4, false},
		{"rotated found left half", []int{4, 5, 6, 7, 0, 1, 2}, 5, true},
		{"rotated found right half", []int{4, 5, 6, 7, 0, 1, 2}, 1, true},
		{"rotated found pivot", []int{4, 5, 6, 7, 0, 1, 2}, 0, true},
		{"rotated not found", []int{4, 5, 6, 7, 0, 1, 2}, 3, false},
		{"rotated found first", []int{4, 5, 6, 7, 0, 1, 2}, 4, true},
		{"rotated found last", []int{4, 5, 6, 7, 0, 1, 2}, 2, true},
		{"duplicates ambiguous boundary found", []int{3, 1, 1}, 3, true},
		{"duplicates ambiguous boundary not found", []int{3, 1, 1}, 2, false},
		{"all duplicates found", []int{2, 2, 2, 2, 2}, 2, true},
		{"all duplicates not found", []int{2, 2, 2, 2, 2}, 3, false},
		{"duplicates rotated found", []int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		{"duplicates rotated not found", []int{2, 5, 6, 0, 0, 1, 2}, 3, false},
		{"single element match", []int{1}, 1, true},
		{"single element no match", []int{1}, 2, false},
	}
	for _, tt := range searchRotatedIITests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchRotatedII(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("SearchRotatedII(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
