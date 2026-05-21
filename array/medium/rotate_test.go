package medium

import (
	"reflect"
	"testing"
)

var rotateTests = []struct {
	name   string
	matrix [][]int
	want   [][]int
}{
	{
		"1x1",
		[][]int{{1}},
		[][]int{{1}},
	},
	{
		"2x2",
		[][]int{{1, 2}, {3, 4}},
		[][]int{{3, 1}, {4, 2}},
	},
	{
		"3x3",
		[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
	},
	{
		"4x4",
		[][]int{
			{5, 1, 9, 11},
			{2, 4, 8, 10},
			{13, 3, 6, 7},
			{15, 14, 12, 16},
		},
		[][]int{
			{15, 13, 2, 5},
			{14, 3, 4, 1},
			{12, 6, 8, 9},
			{16, 7, 10, 11},
		},
	},
	{
		"all same",
		[][]int{{1, 1}, {1, 1}},
		[][]int{{1, 1}, {1, 1}},
	},
	{
		"negative numbers",
		[][]int{{-1, -2}, {-3, -4}},
		[][]int{{-3, -1}, {-4, -2}},
	},
}

func TestRotate(t *testing.T) {
	for _, tt := range rotateTests {
		t.Run(tt.name, func(t *testing.T) {
			matrix := cloneMatrix(tt.matrix)
			Rotate(matrix)
			if !reflect.DeepEqual(matrix, tt.want) {
				t.Errorf("Rotate(%v) = %v, want %v", tt.matrix, matrix, tt.want)
			}
		})
	}
}

func BenchmarkRotate(b *testing.B) {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	for b.Loop() {
		Rotate(cloneMatrix(matrix))
	}
}
