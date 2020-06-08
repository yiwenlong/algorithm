package powx_n

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"normal",
			args{
				2.00000,
				10,
			},
			1024.00000,
		},
		{
			"normal",
			args{
				2.10000,
				3,
			},
			9.26100,
		},
		{
			"normal",
			args{
				0.00001,
				2147483647,
			},
			0.25000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow(%f, %d) = %v, want %v", tt.args.x, tt.args.n, got, tt.want)
			}
		})
	}
}
