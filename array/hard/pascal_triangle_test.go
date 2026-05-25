package hard

import (
	"reflect"
	"testing"
)

var pascalTriangleTests = []struct {
	name string
	n    int
	want [][]int
}{
	{
		"0 rows",
		0,
		[][]int{},
	},
	{
		"1 row",
		1,
		[][]int{{1}},
	},
	{
		"2 rows",
		2,
		[][]int{{1}, {1, 1}},
	},
	{
		"3 rows",
		3,
		[][]int{{1}, {1, 1}, {1, 2, 1}},
	},
	{
		"5 rows",
		5,
		[][]int{
			{1},
			{1, 1},
			{1, 2, 1},
			{1, 3, 3, 1},
			{1, 4, 6, 4, 1},
		},
	},
	{
		"6 rows",
		6,
		[][]int{
			{1},
			{1, 1},
			{1, 2, 1},
			{1, 3, 3, 1},
			{1, 4, 6, 4, 1},
			{1, 5, 10, 10, 5, 1},
		},
	},
}

func TestPascalTriangle(t *testing.T) {
	for _, tt := range pascalTriangleTests {
		t.Run(tt.name, func(t *testing.T) {
			got := PascalTriangle(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PascalTriangle(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkPascalTriangle(b *testing.B) {
	for b.Loop() {
		PascalTriangle(10)
	}
}
