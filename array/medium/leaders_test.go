package medium

import "testing"

var leadersTests = []struct {
	name string
	nums []int
	want string
}{
	{"example", []int{16, 17, 4, 3, 5, 2}, "17 5 2"},
	{"all descending", []int{5, 3, 1}, "5 3 1"},
	{"all ascending", []int{1, 2, 3}, "3"},
	{"single element", []int{7}, "7"},
	{"all same", []int{4, 4, 4}, "4"},
	{"last is largest", []int{1, 2, 3, 10}, "10"},
	{"two elements ascending", []int{1, 2}, "2"},
	{"two elements descending", []int{2, 1}, "2 1"},
	{"negative numbers", []int{-3, -1, -2}, "-1 -2"},
	{"mixed negative positive", []int{-1, 3, -2, 5, 0}, "5 0"},
}

func TestLeaders(t *testing.T) {
	for _, tt := range leadersTests {
		t.Run(tt.name, func(t *testing.T) {
			got := Leaders(tt.nums)
			if got != tt.want {
				t.Errorf("Leaders(%v) = %q, want %q", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkLeaders(b *testing.B) {
	nums := []int{16, 17, 4, 3, 5, 2}
	for b.Loop() {
		Leaders(nums)
	}
}
