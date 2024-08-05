package leetcode

import "testing"

func Test_numberOfSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case-1",
			args: args{"00011"},
			want: 5,
		},
		{
			name: "case-2",
			args: args{"101101"},
			want: 16,
		},
		{
			name: "case-3",
			args: args{"10"},
			want: 2,
		},
		{
			name: "case-4",
			args: args{"011111111"},
			want: 44,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSubstrings(tt.args.s); got != tt.want {
				t.Errorf("numberOfSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
