package utils

import (
	"testing"
)

func TestGetExtFromUrl(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				url: "http://www.foo.com/1.zip",
			},
			want: "zip",
		},
		{
			args: args{
				url: "http://www.foo.com/1.zip?affa=faf",
			},
			want: "zip",
		},
		{
			args: args{
				url: "http://www.foo.com/1.zip?1.a&b",
			},
			want: "zip",
		},
		{
			args: args{
				url: "http://www.foo.com/zip",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetExtFromUrl(tt.args.url); got != tt.want {
				t.Errorf("GetExtFromUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
