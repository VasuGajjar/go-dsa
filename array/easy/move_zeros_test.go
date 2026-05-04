package easy

import (
	"reflect"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"zeros in middle", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"zeros at start", []int{0, 0, 1, 2}, []int{1, 2, 0, 0}},
		{"zeros at end", []int{1, 2, 0, 0}, []int{1, 2, 0, 0}},
		{"all zeros", []int{0, 0, 0}, []int{0, 0, 0}},
		{"no zeros", []int{1, 2, 3}, []int{1, 2, 3}},
		{"single zero", []int{0}, []int{0}},
		{"single non-zero", []int{5}, []int{5}},
		{"alternating", []int{0, 1, 0, 2, 0, 3}, []int{1, 2, 3, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MoveZeros(tt.input)
			if !reflect.DeepEqual(tt.input, tt.want) {
				t.Errorf("MoveZeros result = %v, want %v", tt.input, tt.want)
			}
		})
	}
}
