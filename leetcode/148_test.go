package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				head: makeList([]int{1, 3, 2, 4, 1, 4, 5}),
			},
			want: makeList([]int{1, 1, 2, 3, 4, 4, 5}),
		},
		{
			args: args{
				head: makeList([]int{}),
			},
			want: makeList([]int{}),
		},
		{
			args: args{
				head: makeList([]int{1}),
			},
			want: makeList([]int{1}),
		},
		{
			args: args{
				head: makeList([]int{2, 1, 1}),
			},
			want: makeList([]int{1, 1, 2}),
		},
		{
			args: args{
				head: makeList([]int{1, 1, 1}),
			},
			want: makeList([]int{1, 1, 1}),
		},
		{
			args: args{
				head: makeList([]int{1, 1}),
			},
			want: makeList([]int{1, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
