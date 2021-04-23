package main

import "testing"

func Test_findCombination(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{1, 99, 2, 98, 2, 1, 99, 97, 3, 97},
				k:    100,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCombination(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findCombination() = %v, want %v", got, tt.want)
			}
		})
	}
}
