package add_strings

import "testing"

func Test_addStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"normal",
			args{
				"123",
				"345",
			},
			"468",
		},
		{"normal",
			args{
				"8",
				"345",
			},
			"353",
		},
		{"normal",
			args{
				"345",
				"8",
			},
			"353",
		},
		{"normal",
			args{
				"0",
				"0",
			},
			"0",
		},
		{"normal",
			args{
				"9",
				"1",
			},
			"10",
		},
		{"normal",
			args{
				"1",
				"9",
			},
			"10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings(%s, %s) = %v, want %v", tt.args.num1, tt.args.num2, got, tt.want)
			}
		})
	}
}
