package _sum

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"normal-test",
			args{
				[]int{-1, 0, 1, 2, -1, -4},
			},
			[][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{"three-zero",
			args{
				[]int{0, 0, 0},
			},
			[][]int{{0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
