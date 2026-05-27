package hard

import (
	"fmt"
	"reflect"
	"testing"
)

var triangle = [][]int{
	{1},
	{1, 1},
	{1, 2, 1},
	{1, 3, 3, 1},
	{1, 4, 6, 4, 1},
	{1, 5, 10, 10, 5, 1},
}

func TestPascalTriangle(t *testing.T) {
	for n := 0; n <= len(triangle); n++ {
		t.Run(fmt.Sprintf("%d rows", n), func(t *testing.T) {
			got := PascalTriangle(n)
			want := triangle[:n]
			if n == 0 {
				want = [][]int{}
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("PascalTriangle(%d) = %v, want %v", n, got, want)
			}
		})
	}
}

func TestPascalRow(t *testing.T) {
	for r, row := range triangle {
		r, row := r+1, row
		t.Run(fmt.Sprintf("row %d", r), func(t *testing.T) {
			got := PascalRow(r)
			if !reflect.DeepEqual(got, row) {
				t.Errorf("PascalRow(%d) = %v, want %v", r, got, row)
			}
		})
	}
}

func TestPascalValue(t *testing.T) {
	for r, row := range triangle {
		for c, want := range row {
			r, c, want := r+1, c+1, want
			t.Run(fmt.Sprintf("row%d col%d", r, c), func(t *testing.T) {
				got := PascalValue(r, c)
				if got != want {
					t.Errorf("PascalValue(%d, %d) = %d, want %d", r, c, got, want)
				}
			})
		}
	}
}

func BenchmarkPascalTriangle(b *testing.B) {
	for b.Loop() {
		PascalTriangle(len(triangle))
	}
}

func BenchmarkPascalRow(b *testing.B) {
	for b.Loop() {
		PascalRow(len(triangle))
	}
}

func BenchmarkPascalValue(b *testing.B) {
	for b.Loop() {
		PascalValue(len(triangle), len(triangle)/2)
	}
}
