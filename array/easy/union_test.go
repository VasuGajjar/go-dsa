package easy

import (
	"reflect"
	"testing"
)

func TestUnion(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{"overlapping sorted", []int{1, 2, 3}, []int{2, 3, 4}, []int{1, 2, 3, 4}},
		{"no overlap", []int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"complete overlap", []int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"first is subset", []int{2, 3}, []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{"second is subset", []int{1, 2, 3, 4}, []int{2, 3}, []int{1, 2, 3, 4}},
		{"first empty", []int{}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"second empty", []int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{"both empty", []int{}, []int{}, []int{}},
		{"duplicates in first", []int{1, 1, 2, 3}, []int{2, 4}, []int{1, 2, 3, 4}},
		{"duplicates in second", []int{1, 2}, []int{2, 2, 3, 4}, []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Union(tt.nums1, tt.nums2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
			}
		})
	}
}
