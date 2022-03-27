package utils

import (
	"testing"
)

func TestRemoveBracketsFromString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{str: "12345333"},
			want: "12345333",
		},
		{
			args: args{str: "12345(789)333"},
			want: "12345333",
		},

		{
			args: args{str: "12345 (789) 333"},
			want: "12345  333",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveBracketsFromString(tt.args.str); got != tt.want {
				t.Errorf("RemoveBracketsFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
