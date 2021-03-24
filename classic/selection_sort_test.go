package classic

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {
		// 	args: args{
		// 		nums: []int{1},
		// 	},
		// 	want: []int{1},
		// },
		{
			args: args{
				nums: []int{1, 4, 2, 1, 7, 8, 8},
			},
			want: []int{1, 1, 2, 4, 7, 8, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShellSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
