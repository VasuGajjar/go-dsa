package medium

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"empty", []int{}, 0},
		{"single price", []int{5}, 0},
		{"increasing prices", []int{1, 2, 3, 4, 5}, 4},
		{"decreasing prices", []int{5, 4, 3, 2, 1}, 0},
		{"buy low sell high", []int{7, 1, 5, 3, 6, 4}, 5},
		{"best at end", []int{7, 2, 5, 1, 8}, 7},
		{"best at start and end", []int{1, 8, 2, 9}, 8},
		{"all same price", []int{3, 3, 3}, 0},
		{"two prices profit", []int{1, 5}, 4},
		{"two prices no profit", []int{5, 1}, 0},
		{"dip then rise", []int{6, 1, 3, 2, 4, 7}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxProfit(tt.prices)
			if got != tt.want {
				t.Errorf("MaxProfit(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}
