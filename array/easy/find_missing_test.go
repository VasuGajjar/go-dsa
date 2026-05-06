package easy

import "testing"

var missingTests = []struct {
	name  string
	input []int
	want  int
}{
	{"missing middle", []int{3, 0, 1}, 2},
	{"missing zero", []int{1, 2, 3}, 0},
	{"missing last", []int{0, 1, 2}, 3},
	{"single missing one", []int{0}, 1},
	{"single missing zero", []int{1}, 0},
	{"missing at start", []int{1, 2, 3, 4}, 0},
	{"missing at end", []int{0, 1, 2, 3}, 4},
}

func TestFindMissingWithSum(t *testing.T) {
	for _, tt := range missingTests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMissingWithSum(tt.input)
			if got != tt.want {
				t.Errorf("FindMissingWithSum(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestFindMissingWithXor(t *testing.T) {
	for _, tt := range missingTests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMissingWithXor(tt.input)
			if got != tt.want {
				t.Errorf("FindMissingWithXor(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

var largeMissingInput = func() []int {
	n := 10000
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i
	}
	// missing number is n (last)
	return nums
}()

func BenchmarkFindMissingWithSum(b *testing.B) {
	for b.Loop() {
		FindMissingWithSum(largeMissingInput)
	}
}

func BenchmarkFindMissingWithXor(b *testing.B) {
	for b.Loop() {
		FindMissingWithXor(largeMissingInput)
	}
}
