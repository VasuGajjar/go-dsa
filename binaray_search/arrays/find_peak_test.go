package arrays

import "testing"

func TestFindPeak(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int // valid peak indices (multiple possible)
	}{
		{"single element", []int{1}, []int{0}},
		{"peak at start", []int{5, 3, 1}, []int{0}},
		{"peak at end", []int{1, 3, 5}, []int{2}},
		{"peak in middle", []int{1, 3, 2}, []int{1}},
		{"peak first of two", []int{2, 1}, []int{0}},
		{"peak second of two", []int{1, 2}, []int{1}},
		{"multiple peaks left", []int{1, 5, 3, 4, 2}, []int{1, 3}},
		{"multiple peaks right", []int{3, 1, 4, 2, 5}, []int{0, 2, 4}},
		{"strictly increasing", []int{1, 2, 3, 4, 5}, []int{4}},
		{"strictly decreasing", []int{5, 4, 3, 2, 1}, []int{0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindPeak(tt.nums)
			valid := false
			for _, w := range tt.want {
				if got == w {
					valid = true
					break
				}
			}
			if !valid {
				t.Errorf("FindPeak(%v) = %d, want one of %v", tt.nums, got, tt.want)
			}
		})
	}
}
