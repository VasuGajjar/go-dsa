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

var pascalRowTests = []struct {
	name string
	r    int
	want []int
}{
	{"row 1", 1, []int{1}},
	{"row 2", 2, []int{1, 1}},
	{"row 3", 3, []int{1, 2, 1}},
	{"row 4", 4, []int{1, 3, 3, 1}},
	{"row 5", 5, []int{1, 4, 6, 4, 1}},
	{"row 6", 6, []int{1, 5, 10, 10, 5, 1}},
	{"row 7", 7, []int{1, 6, 15, 20, 15, 6, 1}},
}

func TestPascalRow(t *testing.T) {
	for _, tt := range pascalRowTests {
		t.Run(tt.name, func(t *testing.T) {
			got := PascalRow(tt.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PascalRow(%d) = %v, want %v", tt.r, got, tt.want)
			}
		})
	}
}
