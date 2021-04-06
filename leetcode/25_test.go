package leetcode

import (
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				head: makeList([]int{1, 2, 3, 4}),
				k:    2,
			},
			want: makeList([]int{2, 1, 4, 3}),
		},
		{
			args: args{
				head: makeList([]int{1, 2}),
				k:    2,
			},
			want: makeList([]int{2, 1}),
		},
		{
			args: args{
				head: makeList([]int{1}),
				k:    2,
			},
			want: makeList([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
