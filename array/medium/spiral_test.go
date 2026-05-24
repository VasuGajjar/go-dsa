package medium

import (
	"reflect"
	"testing"
)

var spiralTests = []struct {
	name   string
	matrix [][]int
	want   []int
}{
	{
		"1x1",
		[][]int{{1}},
		[]int{1},
	},
	{
		"2x2",
		[][]int{{1, 2}, {3, 4}},
		[]int{1, 2, 4, 3},
	},
	{
		"3x3",
		[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
	},
	{
		"4x4",
		[][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		},
		[]int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10},
	},
	{
		"1x4 single row",
		[][]int{{1, 2, 3, 4}},
		[]int{1, 2, 3, 4},
	},
	{
		"4x1 single column",
		[][]int{{1}, {2}, {3}, {4}},
		[]int{1, 2, 3, 4},
	},
	{
		"2x3",
		[][]int{{1, 2, 3}, {4, 5, 6}},
		[]int{1, 2, 3, 6, 5, 4},
	},
	{
		"3x2",
		[][]int{{1, 2}, {3, 4}, {5, 6}},
		[]int{1, 2, 4, 6, 5, 3},
	},
	{
		"negative numbers",
		[][]int{{-1, -2}, {-3, -4}},
		[]int{-1, -2, -4, -3},
	},
}

func TestSpiralOrder(t *testing.T) {
	for _, tt := range spiralTests {
		t.Run(tt.name, func(t *testing.T) {
			got := SpiralOrder(tt.matrix)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpiralOrder(%v) = %v, want %v", tt.matrix, got, tt.want)
			}
		})
	}
}

func BenchmarkSpiralOrder(b *testing.B) {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	for b.Loop() {
		SpiralOrder(matrix)
	}
}
