package answer

import "testing"

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		name  string
		piles []int
		h     int
		want  int
	}{
		{"empty piles", []int{}, 5, -1},
		{"nil piles", nil, 5, -1},
		{"fewer hours than piles", []int{3, 6, 7, 11}, 3, -1},
		{"exact hours per pile", []int{30, 11, 23, 4, 20}, 5, 30},
		{"one spare hour", []int{30, 11, 23, 4, 20}, 6, 23},
		{"plenty of hours", []int{3, 6, 7, 11}, 8, 4},
		{"single pile, lots of time", []int{312884470}, 968709470, 1},
		{"single pile, two hours", []int{1000000000}, 2, 500000000},
		{"single pile, odd split", []int{7}, 2, 4},
		{"all ones, hours equal piles", []int{1, 1, 1, 1}, 4, 1},
		{"all ones, extra hours", []int{1, 1, 1, 1}, 10, 1},
		{"equal piles", []int{5, 5, 5}, 6, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinEatingSpeed(tt.piles, tt.h)
			if got != tt.want {
				t.Errorf("MinEatingSpeed(%v, %d) = %d, want %d", tt.piles, tt.h, got, tt.want)
			}
		})
	}
}