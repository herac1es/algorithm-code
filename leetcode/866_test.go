package leetcode

import "testing"

func Test_primePalindrome(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	args: args{
		// 		N: 6,
		// 	},
		// 	want: 7,
		// },
		// {
		// 	args: args{
		// 		N: 8,
		// 	},
		// 	want: 11,
		// },
		{
			args: args{
				N: 9602070,
			},
			want: 101,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := primePalindrome(tt.args.N); got != tt.want {
				t.Errorf("primePalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
