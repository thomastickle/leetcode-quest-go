package main

import (
	"slices"
	"testing"
)

func Test_shuffle(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		n    int
		want []int
	}{
		{"Case 1", []int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
		{"Case 2", []int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{"Case 3", []int{1, 1, 2, 2}, 2, []int{1, 2, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shuffle(tt.nums, tt.n)
			// TODO: update the condition below to compare got with tt.want.
			if !slices.Equal(tt.want, got) {
				t.Errorf("shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}
