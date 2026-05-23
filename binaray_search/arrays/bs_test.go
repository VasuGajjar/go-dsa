package arrays

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		nums     []int
		val      int
		expected int
	}{
		{[]int{1, 3, 5, 7, 9}, 5, 2},
		{[]int{1, 3, 5, 7, 9}, 1, 0},
		{[]int{1, 3, 5, 7, 9}, 9, 4},
		{[]int{1, 3, 5, 7, 9}, 4, -1},
		{[]int{1}, 1, 0},
		{[]int{1}, 2, -1},
		{[]int{}, 1, -1},
	}

	for _, tt := range tests {
		got := BinarySearch(tt.nums, tt.val)
		if got != tt.expected {
			t.Errorf("BinarySearch(%v, %d) = %d, want %d", tt.nums, tt.val, got, tt.expected)
		}
	}
}
