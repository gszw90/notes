package test

import (
	"github.com/gszw90/notes/algorithm"
	"testing"
)

func Test_maximumLength(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{s: "aaaa"},
			want: 2,
		},
		{
			name: "test2",
			args: args{s: "abcdef"},
			want: -1,
		},
		{
			name: "test3",
			args: args{s: "abcaba"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := algorithm.MaximumLength(tt.args.s); got != tt.want {
				t.Errorf("maximumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
