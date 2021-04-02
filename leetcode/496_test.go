package leetcode

import (
	"reflect"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums1: []int{4, 1, 2},
				nums2: []int{1, 3, 4, 2},
			},
			want: []int{-1, 3, -1},
		},
		{
			args: args{
				nums1: []int{2, 4},
				nums2: []int{1, 2, 3, 4},
			},
			want: []int{3, -1},
		},
		{
			args: args{
				nums1: []int{1},
				nums2: []int{1},
			},
			want: []int{-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextGreaterElement(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
