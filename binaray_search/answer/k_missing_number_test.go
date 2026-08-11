package answer

import "testing"

func TestFindKthMissing(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		k    int
		want int
	}{
		{"empty array", []int{}, 5, 5},
		{"nil array", nil, 1, 1},
		{"first element beyond k", []int{5, 6}, 3, 3},
		{"first element equals k", []int{2}, 2, 3},
		{"no gaps, missing after array", []int{1, 2, 3}, 1, 4},
		{"no gaps, far beyond array", []int{1, 2, 3}, 5, 8},
		{"gap in middle", []int{2, 3, 4, 7, 11}, 5, 9},
		{"first missing before array", []int{2, 3, 4, 7, 11}, 1, 1},
		{"missing inside gap", []int{2, 3, 4, 7, 11}, 2, 5},
		{"last missing inside array range", []int{2, 3, 4, 7, 11}, 6, 10},
		{"missing past last element", []int{2, 3, 4, 7, 11}, 7, 12},
		{"single element with leading gap", []int{5}, 4, 4},
		{"single element, missing after it", []int{5}, 5, 6},
		{"large gap", []int{1, 100}, 50, 51},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindKthMissing(tt.arr, tt.k)
			if got != tt.want {
				t.Errorf("FindKthMissing(%v, %d) = %d, want %d", tt.arr, tt.k, got, tt.want)
			}
		})
	}
}