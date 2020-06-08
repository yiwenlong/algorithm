package main

import (
	"testing"
)

var board1 = [][]byte{
	{'E', 'E', 'E', 'E', 'E'},
	{'E', 'E', 'M', 'E', 'E'},
	{'E', 'E', 'E', 'E', 'E'},
	{'E', 'E', 'E', 'E', 'E'},
}

var dstBoard1 = [][]byte{
	{'B', 'E', 'E', 'E', 'B'},
	{'B', 'E', 'M', 'E', 'B'},
	{'B', 'E', 'E', 'E', 'B'},
	{'B', 'B', 'B', 'B', 'B'},
}

var board2 = [][]byte{
	{'B', '1', 'E', '1', 'B'},
	{'B', '1', 'M', '1', 'B'},
	{'B', '1', '1', '1', 'B'},
	{'B', 'B', 'B', 'B', 'B'},
}

var board3 = [][]byte{
	{'B', 'B', 'B', 'B', 'B', 'B', '1', 'E'},
	{'B', '1', '1', '1', 'B', 'B', '1', 'M'},
	{'1', '2', 'M', '1', 'B', 'B', '1', '1'},
	{'M', '2', '1', '1', 'B', 'B', 'B', 'B'},
	{'1', '1', 'B', 'B', 'B', 'B', 'B', 'B'},
	{'B', 'B', 'B', 'B', 'B', 'B', 'B', 'B'},
	{'B', '1', '2', '2', '1', 'B', 'B', 'B'},
	{'B', '1', 'M', 'M', '1', 'B', 'B', 'B'},
}

func Test_mineAroundCount(t *testing.T) {
	type args struct {
		board    [][]byte
		position []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{board1, []int{3, 0}}, 0},
		{"test2", args{board1, []int{0, 1}}, 1},
		{"test3", args{board1, []int{0, 2}}, 1},
		{"test4", args{board1, []int{0, 3}}, 1},
		{"test5", args{board1, []int{0, 4}}, 0},
		{"test6", args{board1, []int{1, 0}}, 0},
		{"test7", args{board1, []int{1, 1}}, 1},
		{"test8", args{board1, []int{1, 3}}, 1},
		{"test9", args{board1, []int{2, 1}}, 1},
		{"test10", args{board1, []int{2, 3}}, 1},
		{"test11", args{board3, []int{0, 7}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mineAroundCount(tt.args.board, tt.args.position); got != tt.want {
				t.Errorf("mineAroundCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
