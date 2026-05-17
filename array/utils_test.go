package array

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "odd length",
			nums:     []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "even length",
			nums:     []int{1, 2, 3, 4},
			expected: []int{4, 3, 2, 1},
		},
		{
			name:     "single element",
			nums:     []int{7},
			expected: []int{7},
		},
		{
			name:     "two elements",
			nums:     []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "already reversed",
			nums:     []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "all same elements",
			nums:     []int{3, 3, 3},
			expected: []int{3, 3, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.nums))
			copy(input, tt.nums)
			Reverse(input)
			if !reflect.DeepEqual(input, tt.expected) {
				t.Errorf("reverse(%v) = %v, want %v", tt.nums, input, tt.expected)
			}
		})
	}
}
