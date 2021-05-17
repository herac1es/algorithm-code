package classic

import (
	"math/rand"
	"reflect"
	"sort"
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
		{
			args: args{
				nums: []int{1},
			},
			want: []int{1},
		},
		{
			args: args{
				nums: []int{1, 3, 2, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			args: args{
				nums: []int{-1, 2, 2, 2, 1, -1},
			},
			want: []int{-1, -1, 1, 2, 2, 2},
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

type ss []int

func (s ss) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s ss) Len() int {
	return len(s)
}
func (s ss) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func BenchmarkSort(b *testing.B) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				nums: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 23, 2, 2, 42, 1, 2, 12323, 232, 2, 323, 122, 3, 2322, 32, 32, 2, 1, 3, 5, 6, 7, 464, 6453, 3, 3434, 3, 6353, 4564, 657, 53, 343, 43, 343, 4234, 34, 3434, 34343, 434, 343, 34, 342, 34, 334, 354, 4, 6464545},
			},
		},
	}
	for _, tt := range tests {
		rand.Shuffle(len(tt.args.nums), func(i, j int) {
			tt.args.nums[i], tt.args.nums[j] = tt.args.nums[j], tt.args.nums[i]
		})
		b.Run("test:std", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sort.Sort(ss(tt.args.nums))
			}
		})
		b.Run("test:quick", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				QuickSort(tt.args.nums)
			}
		})
		b.Run("test:mergeSort", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MergeSort(tt.args.nums)
			}
		})
		b.Run("test:shell sort", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ShellSort(tt.args.nums)
			}
		})
	}
}
