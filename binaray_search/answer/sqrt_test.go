package answer

import "testing"

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

func TestNthRoot(t *testing.T) {
	tests := []struct {
		name string
		num  int
		k    int
		want int
	}{
		{
			name: "perfect square",
			num:  16,
			k:    2,
			want: 4,
		},
		{
			name: "perfect cube",
			num:  27,
			k:    3,
			want: 3,
		},
		{
			name: "perfect fourth root",
			num:  81,
			k:    4,
			want: 3,
		},
		{
			name: "non perfect square returns floor",
			num:  20,
			k:    2,
			want: 4,
		},
		{
			name: "non perfect cube returns floor",
			num:  30,
			k:    3,
			want: 3,
		},
		{
			name: "zero",
			num:  0,
			k:    2,
			want: 0,
		},
		{
			name: "one square root",
			num:  1,
			k:    2,
			want: 1,
		},
		{
			name: "one cube root",
			num:  1,
			k:    3,
			want: 1,
		},
		{
			name: "root equals number",
			num:  8,
			k:    1,
			want: 8,
		},
		{
			name: "large perfect square",
			num:  1000000,
			k:    2,
			want: 1000,
		},
		{
			name: "large non perfect square",
			num:  1000001,
			k:    2,
			want: 1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NthRoot(tt.num, tt.k)
			if got != tt.want {
				t.Errorf("NthRoot(%d, %d) = %d, want %d",
					tt.num, tt.k, got, tt.want)
			}
		})
	}
}
