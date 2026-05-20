package medium

import (
	"reflect"
	"testing"
)

func cloneMatrix(m [][]int) [][]int {
	c := make([][]int, len(m))
	for i := range m {
		c[i] = make([]int, len(m[i]))
		copy(c[i], m[i])
	}
	return c
}

var zeroMatrixTests = []struct {
	name   string
	matrix [][]int
	want   [][]int
}{
	{
		"no zeros",
		[][]int{{1, 2}, {3, 4}},
		[][]int{{1, 2}, {3, 4}},
	},
	{
		"zero in middle",
		[][]int{{1, 2, 3}, {4, 0, 6}, {7, 8, 9}},
		[][]int{{1, 0, 3}, {0, 0, 0}, {7, 0, 9}},
	},
	{
		"zero in first row",
		[][]int{{1, 0, 3}, {4, 5, 6}, {7, 8, 9}},
		[][]int{{0, 0, 0}, {4, 0, 6}, {7, 0, 9}},
	},
	{
		"zero in first col",
		[][]int{{1, 2, 3}, {0, 5, 6}, {7, 8, 9}},
		[][]int{{0, 2, 3}, {0, 0, 0}, {0, 8, 9}},
	},
	{
		"zero at top-left",
		[][]int{{0, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		[][]int{{0, 0, 0}, {0, 5, 6}, {0, 8, 9}},
	},
	{
		"multiple zeros",
		[][]int{{1, 0, 3}, {4, 5, 6}, {7, 8, 0}},
		[][]int{{0, 0, 0}, {4, 0, 0}, {0, 0, 0}},
	},
	{
		"all zeros",
		[][]int{{0, 0}, {0, 0}},
		[][]int{{0, 0}, {0, 0}},
	},
	{
		"single element zero",
		[][]int{{0}},
		[][]int{{0}},
	},
	{
		"single element non-zero",
		[][]int{{5}},
		[][]int{{5}},
	},
	{
		"single row",
		[][]int{{1, 0, 3}},
		[][]int{{0, 0, 0}},
	},
	{
		"single col",
		[][]int{{1}, {0}, {3}},
		[][]int{{0}, {0}, {0}},
	},
	{
		"zero in last row last col",
		[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}},
		[][]int{{1, 2, 0}, {4, 5, 0}, {0, 0, 0}},
	},
}

func TestZeroMatrix(t *testing.T) {
	for _, tt := range zeroMatrixTests {
		t.Run(tt.name, func(t *testing.T) {
			matrix := cloneMatrix(tt.matrix)
			ZeroMatrix(matrix)
			if !reflect.DeepEqual(matrix, tt.want) {
				t.Errorf("ZeroMatrix(%v) = %v, want %v", tt.matrix, matrix, tt.want)
			}
		})
	}
}

func BenchmarkZeroMatrix(b *testing.B) {
	matrix := [][]int{
		{1, 0, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 0, 12},
		{13, 14, 15, 16},
	}
	for b.Loop() {
		ZeroMatrix(cloneMatrix(matrix))
	}
}
