package main

import "testing"

func TestQuickSort(t *testing.T) {
	type args struct {
		arr   []int
		begin int
		end   int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{[]int{2, 1}, 0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("res: %v\n", tt.args.arr)
		})
	}
}
