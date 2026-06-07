package answer

import "testing"

func TestSquareRoot(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{"0", 0, 0},
		{"1", 1, 1},
		{"perfect square 4", 4, 2},
		{"perfect square 9", 9, 3},
		{"perfect square 16", 16, 4},
		{"perfect square 25", 25, 5},
		{"perfect square 100", 100, 10},
		{"floor sqrt 2", 2, 1},
		{"floor sqrt 3", 3, 1},
		{"floor sqrt 8", 8, 2},
		{"floor sqrt 10", 10, 3},
		{"floor sqrt 26", 26, 5},
		{"floor sqrt 99", 99, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SquareRoot(tt.num)
			if got != tt.want {
				t.Errorf("SquareRoot(%d) = %d, want %d", tt.num, got, tt.want)
			}
		})
	}
}

func TestPow(t *testing.T) {
	tests := []struct {
		num, exp, want int
	}{
		{2, 0, 1},
		{2, 1, 2},
		{2, 10, 1024},
		{3, 3, 27},
		{5, 4, 625},
		{1, 100, 1},
		{10, 3, 1000},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := Pow(tt.num, tt.exp)
			if got != tt.want {
				t.Errorf("Pow(%d, %d) = %d, want %d", tt.num, tt.exp, got, tt.want)
			}
		})
	}
}
