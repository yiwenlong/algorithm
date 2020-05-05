package minimum_genetic_mutation

import "testing"

func Test_minMutation(t *testing.T) {
	type args struct {
		start string
		end   string
		bank  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"testcase-1",
			args{
				"AACCGGTT",
				"AACCGGTA",
				[]string{"AACCGGTA"},
			},
			1,
		},

		{
			"testcase-2",
			args{
				"AACCGGTT",
				"AAACGGTA",
				[]string{"AACCGGTA", "AACCGCTA", "AAACGGTA"},
			},
			2,
		},

		{
			"testcase-3",
			args{
				"AAAAACCC",
				"AACCCCCC",
				[]string{"AAAACCCC", "AAACCCCC", "AACCCCCC"},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMutation(tt.args.start, tt.args.end, tt.args.bank); got != tt.want {
				t.Errorf("minMutation() = %v, want %v", got, tt.want)
			}
		})
	}
}