package sort

var tests = []struct {
	name  string
	input []int
	want  []int
}{
	{"sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	{"reverse", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	{"unsorted", []int{3, 1, 4, 1, 5, 9, 2, 6}, []int{1, 1, 2, 3, 4, 5, 6, 9}},
	{"single", []int{42}, []int{42}},
	{"empty", []int{}, []int{}},
	{"duplicates", []int{3, 3, 1, 1, 2}, []int{1, 1, 2, 3, 3}},
	{"negatives", []int{-3, 0, -1, 2, -5}, []int{-5, -3, -1, 0, 2}},
}
