package medium

import (
	"reflect"
	"testing"
)

func TestRearrangeArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "basic example",
			nums: []int{3, 1, -2, -5, 2, -4},
			want: []int{3, -2, 1, -5, 2, -4},
		},
		{
			name: "starts with negative",
			nums: []int{-1, 1},
			want: []int{1, -1},
		},
		{
			name: "includes zero",
			nums: []int{0, -1, 2, -3},
			want: []int{0, -1, 2, -3},
		},
		{
			name: "single positive and negative",
			nums: []int{5, -7},
			want: []int{5, -7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RearrangeArray(tt.nums)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RearrangeArray(%v) = %v, want %v",
					tt.nums, got, tt.want)
			}
		})
	}
}
