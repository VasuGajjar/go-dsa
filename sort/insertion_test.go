package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestInsertion(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Insertion(tt.input)
			for i, v := range tt.input {
				if v != tt.want[i] {
					t.Errorf("got %v, want %v", tt.input, tt.want)
					break
				}
			}
		})
	}
}

func BenchmarkInsertion(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	for _, n := range sizes {
		data := rand.Perm(n)
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for b.Loop() {
				input := make([]int, n)
				copy(input, data)
				Insertion(input)
			}
		})
	}
}
