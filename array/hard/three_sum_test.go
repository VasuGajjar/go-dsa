package hard

import (
	"reflect"
	"testing"
)

var threeSumTests = []struct {
	name string
	nums []int
	want [][]int
}{
	{
		"example",
		[]int{-1, 0, 1, 2, -1, -4},
		[][]int{{-1, -1, 2}, {-1, 0, 1}},
	},
	{
		"all zeros",
		[]int{0, 0, 0, 0},
		[][]int{{0, 0, 0}},
	},
	{
		"no triplet",
		[]int{1, 2, -2, -1},
		[][]int{},
	},
	{
		"fewer than three elements",
		[]int{1, 2},
		[][]int{},
	},
	{
		"empty",
		[]int{},
		[][]int{},
	},
	{
		"duplicates produce single triplet",
		[]int{0, 0, 0},
		[][]int{{0, 0, 0}},
	},
	{
		"negative and positive mix",
		[]int{-2, 0, 1, 1, 2},
		[][]int{{-2, 0, 2}, {-2, 1, 1}},
	},
	{
		"many duplicates yield unique triplets only",
		[]int{-4, -2, -2, -2, 0, 0, 0, 2, 2, 2, 4},
		[][]int{{-4, 0, 4}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}, {0, 0, 0}},
	},
}

func TestThreeSum(t *testing.T) {
	for _, tt := range threeSumTests {
		t.Run(tt.name, func(t *testing.T) {
			got := ThreeSum(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThreeSum(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkThreeSum(b *testing.B) {
	nums := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 4}
	for b.Loop() {
		ThreeSum(nums)
	}
}
