package leetcode

import "testing"

func Test_fourSumCount(t *testing.T) {
	type args struct {
		A []int
		B []int
		C []int
		D []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				A: []int{-1, -1},
				B: []int{-1, 1},
				C: []int{-1, 1},
				D: []int{1, -1},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSumCount(tt.args.A, tt.args.B, tt.args.C, tt.args.D); got != tt.want {
				t.Errorf("fourSumCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
