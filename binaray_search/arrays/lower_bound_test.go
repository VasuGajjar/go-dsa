package arrays

import "testing"

var lowerBoundTests = []struct {
	name string
	nums []int
	k    int
	want int
}{
	{"exact match first", []int{1, 3, 5, 7, 9}, 1, 0},
	{"exact match mid", []int{1, 3, 5, 7, 9}, 5, 2},
	{"exact match last", []int{1, 3, 5, 7, 9}, 9, 4},
	{"between elements", []int{1, 3, 5, 7, 9}, 4, 2},
	{"less than all", []int{1, 3, 5, 7, 9}, 0, 0},
	{"greater than all", []int{1, 3, 5, 7, 9}, 10, 5},
	{"single element match", []int{5}, 5, 0},
	{"single element less", []int{5}, 3, 0},
	{"single element greater", []int{5}, 7, 1},
	{"duplicates at target", []int{1, 3, 3, 3, 5}, 3, 1},
	{"duplicates before target", []int{1, 1, 1, 3, 5}, 2, 3},
}

func TestLowerBound(t *testing.T) {
	for _, tt := range lowerBoundTests {
		t.Run(tt.name, func(t *testing.T) {
			got := LowerBound(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("LowerBound(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
