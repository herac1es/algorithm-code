package leetcode

import (
	"reflect"
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				head: makeList([]int{1, 2, 6, 3, 4, 5, 6}),
				val:  6,
			},
			want: makeList([]int{1, 2, 3, 4, 5}),
		},
		{
			args: args{
				head: makeList([]int{}),
				val:  6,
			},
			want: makeList([]int{}),
		},
		{
			args: args{
				head: makeList([]int{1}),
				val:  1,
			},
			want: makeList([]int{}),
		},
		{
			args: args{
				head: makeList([]int{1, 2, 1, 1}),
				val:  1,
			},
			want: makeList([]int{2}),
		},
		{
			args: args{
				head: makeList([]int{1, 1, 1, 1}),
				val:  1,
			},
			want: makeList([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
