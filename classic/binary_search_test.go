package classic

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums:   []int{},
				target: 0,
			},
			want: -1,
		},
		{
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: 0,
		},
		{
			args: args{
				nums:   []int{1},
				target: 2,
			},
			want: -1,
		},
		{
			args: args{
				nums:   []int{1, 2, 4, 5, 6},
				target: 5,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
