package easy

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"single at end", []int{2, 2, 1}, 1},
		{"single at start", []int{4, 1, 2, 1, 2}, 4},
		{"single in middle", []int{1, 3, 1}, 3},
		{"only one element", []int{7}, 7},
		{"larger input", []int{4, 1, 2, 9, 1, 2, 4}, 9},
		{"negative single", []int{-1, 2, 2}, -1},
		{"negative duplicate single positive", []int{-3, -3, 5}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SingleNumberWithMap(tt.input)
			if got != tt.want {
				t.Errorf("SingleNumberWithMap(%v) = %d, want %d", tt.input, got, tt.want)
			}

			got = SingleNumberWithXor(tt.input)
			if got != tt.want {
				t.Errorf("SingleNumberWithXor(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

var largeSingleInput = func() []int {
	n := 10000
	nums := make([]int, 2*n+1)
	for i := range n {
		nums[2*i] = i + 1
		nums[2*i+1] = i + 1
	}
	nums[2*n] = n + 1 // single number at end
	return nums
}()

func BenchmarkSingleNumberWithMap(b *testing.B) {
	for b.Loop() {
		SingleNumberWithMap(largeSingleInput)
	}
}

func BenchmarkSingleNumberWithXor(b *testing.B) {
	for b.Loop() {
		SingleNumberWithXor(largeSingleInput)
	}
}
