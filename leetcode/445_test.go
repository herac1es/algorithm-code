package leetcode

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbersII(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				l1: makeList([]int{9}),
				l2: makeList([]int{9, 9, 9}),
			},
			want: makeList([]int{1, 0, 0, 8}),
		},

		{
			args: args{
				l1: makeList([]int{7, 2, 4, 3}),
				l2: makeList([]int{5, 6, 4}),
			},
			want: makeList([]int{7, 8, 0, 7}),
		},
		{
			args: args{
				l1: makeList([]int{7}),
				l2: makeList([]int{5}),
			},
			want: makeList([]int{1, 2}),
		},
		{
			args: args{
				l1: makeList([]int{1}),
				l2: makeList([]int{0}),
			},
			want: makeList([]int{1}),
		},
		{
			args: args{
				l1: makeList([]int{0}),
				l2: makeList([]int{0}),
			},
			want: makeList([]int{0}),
		},
		{
			args: args{
				l1: makeList([]int{1}),
				l2: makeList([]int{9, 9}),
			},
			want: makeList([]int{1, 0, 0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbersII(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbersII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func makeList(nums []int) *ListNode {
	head := new(ListNode)
	cur := head
	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return head.Next
}
