package arrays

import "testing"

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

func TestSearchRotated(t *testing.T) {
	for _, tt := range searchRotatedTests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchRotated(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("SearchRotated(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
