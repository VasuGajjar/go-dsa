package hard

import (
	"slices"
	"testing"
)

var majorityElementTests = []struct {
	name string
	nums []int
	want []int
}{
	{
		"single element",
		[]int{1},
		[]int{1},
	},
	{
		"all same",
		[]int{2, 2, 2},
		[]int{2},
	},
	{
		"one majority",
		[]int{3, 2, 3},
		[]int{3},
	},
	{
		"two majority",
		[]int{1, 1, 1, 2, 2, 2, 3},
		[]int{1, 2},
	},
	{
		"no majority",
		[]int{1, 2, 3, 4, 5, 6},
		[]int{},
	},
	{
		"majority at start",
		[]int{1, 1, 1, 1, 2, 3, 4},
		[]int{1},
	},
	{
		"majority at end",
		[]int{2, 3, 4, 1, 1, 1, 1},
		[]int{1},
	},
	{
		"negative numbers",
		[]int{-1, -1, -1, 2, 2, 3},
		[]int{-1},
	},
	{
		"two majority unsorted",
		[]int{1, 2, 1, 2, 1, 2, 3},
		[]int{1, 2},
	},
}

func TestMajorityElement(t *testing.T) {
	for _, tt := range majorityElementTests {
		t.Run(tt.name, func(t *testing.T) {
			got := MajorityElement(tt.nums)
			slices.Sort(got)
			slices.Sort(tt.want)
			if len(got) != len(tt.want) {
				t.Errorf("MajorityElement(%v) = %v, want %v", tt.nums, got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("MajorityElement(%v) = %v, want %v", tt.nums, got, tt.want)
					return
				}
			}
		})
	}
}

func BenchmarkMajorityElement(b *testing.B) {
	nums := []int{1, 1, 1, 2, 2, 2, 3, 4, 5, 6}
	for b.Loop() {
		MajorityElement(nums)
	}
}
