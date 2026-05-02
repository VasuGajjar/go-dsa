package easy

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"no duplicates", []int{1, 2, 3}, []int{1, 2, 3}},
		{"all duplicates", []int{1, 1, 1}, []int{1, 0, 0}},
		{"some duplicates", []int{1, 1, 2, 3, 3}, []int{1, 2, 3, 0, 0}},
		{"single element", []int{5}, []int{5}},
		{"empty", []int{}, []int{}},
		{"two same", []int{4, 4}, []int{4, 0}},
		{"negatives", []int{-3, -3, -1, 0, 0, 2}, []int{-3, -1, 0, 2, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RemoveDuplicates(tt.input)
			for i, v := range tt.input {
				if v != tt.want[i] {
					t.Errorf("RemoveDuplicates => got %v, want %v", tt.input, tt.want)
					break
				}
			}
		})
	}
}
