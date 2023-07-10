package main

import "testing"

func Test_findMaxSpace(t *testing.T) {
	tests := []struct {
		name   string
		cows   int
		stalls []int
		want   int
	}{
		{
			name:   "max space with 2 cows and 2 stalls",
			cows:   2,
			stalls: []int{1, 5},
			want:   4,
		},
		{
			name:   "largest max space with 3 cows and 5 stalls",
			cows:   3,
			stalls: []int{1, 2, 8, 4, 9},
			want:   3,
		},
		{
			name:   "largest max space with 4 cows and 6 stalls",
			cows:   4,
			stalls: []int{1, 37, 11, 23, 32},
			want:   10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxSpace(tt.cows, tt.stalls); got != tt.want {
				t.Errorf("findMaxSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
