package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestQuick(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Quick(tt.input)
			for i, v := range tt.input {
				if v != tt.want[i] {
					t.Errorf("got %v, want %v", tt.input, tt.want)
					break
				}
			}
		})
	}
}

func TestPartition(t *testing.T) {
	var tests = []struct {
		name  string
		input []int
		want  int
	}{
		{"sorted", []int{1, 2, 3, 4, 5}, 0},
		{"reverse", []int{5, 4, 3, 2, 1}, 4},
		{"unsorted", []int{3, 1, 4, 1, 5, 9, 2, 6}, 3},
		{"single", []int{42}, 0},
		{"empty", []int{}, 0},
		{"duplicates", []int{3, 3, 1, 1, 2}, 4},
		{"negatives", []int{-3, 0, -1, 2, -5}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val := partition(tt.input)
			if val != tt.want {
				t.Errorf("%v => got %v want %v", tt.input, val, tt.want)
			}
		})
	}
}

func BenchmarkQuick(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	for _, n := range sizes {
		data := rand.Perm(n)
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for b.Loop() {
				input := make([]int, n)
				copy(input, data)
				Quick(input)
			}
		})
	}
}
