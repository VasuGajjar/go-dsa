package medium

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"single element", []int{1}, 1},
		{"two same elements", []int{3, 3}, 3},
		{"majority at start", []int{3, 3, 1}, 3},
		{"majority at end", []int{1, 3, 3}, 3},
		{"majority in middle", []int{1, 2, 2, 2, 1}, 2},
		{"all same", []int{5, 5, 5, 5}, 5},
		{"majority by one", []int{2, 2, 1, 1, 2}, 2},
		{"larger input", []int{6, 5, 5, 6, 6, 6, 3}, 6},
		{"negative majority", []int{-1, -1, 2}, -1},
		{"empty slice", []int{}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MajorityElement(tt.nums)
			if got != tt.want {
				t.Errorf("MajorityElement(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
