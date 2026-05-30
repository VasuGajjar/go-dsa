package arrays

import "testing"

var upperBoundTests = []struct {
	name string
	nums []int
	k    int
	want int
}{
	{"exact match first", []int{1, 3, 5, 7, 9}, 1, 1},
	{"exact match mid", []int{1, 3, 5, 7, 9}, 5, 3},
	{"exact match last", []int{1, 3, 5, 7, 9}, 9, 5},
	{"between elements", []int{1, 3, 5, 7, 9}, 4, 2},
	{"less than all", []int{1, 3, 5, 7, 9}, 0, 0},
	{"greater than all", []int{1, 3, 5, 7, 9}, 10, 5},
	{"single element match", []int{5}, 5, 1},
	{"single element less", []int{5}, 3, 0},
	{"single element greater", []int{5}, 7, 1},
	{"duplicates at target", []int{1, 3, 3, 3, 5}, 3, 4},
	{"duplicates before target", []int{1, 1, 1, 3, 5}, 1, 3},
}

func TestUpperBound(t *testing.T) {
	for _, tt := range upperBoundTests {
		t.Run(tt.name, func(t *testing.T) {
			got := UpperBound(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("UpperBound(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
