package sliding_window_maximum

import (
	"reflect"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"normal",
			args{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3},
			[]int{3, 3, 5, 5, 6, 7},
		},
		{
			"normal",
			args{[]int{1, 3, 1, 2, 0, 5}, 3},
			[]int{3, 3, 2, 5},
		},
		{
			"normal",
			args{[]int{9, 10, 9, -7, -4, -8, 2, -6}, 5},
			[]int{10, 10, 9, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow(%v, %d) = %v, want %v", tt.args.nums, tt.args.k, got, tt.want)
			}
		})
	}
}
