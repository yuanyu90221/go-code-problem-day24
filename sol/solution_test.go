package sol

import "testing"

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "a2[B2[c]]",
			args: args{
				s: "a2[B2[c]]",
			},
			want: "aBccBcc",
		},
		{
			name: "a2[bc]12[d]",
			args: args{
				s: "a2[bc]12[d]",
			},
			want: "abcbcdddddddddddd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
