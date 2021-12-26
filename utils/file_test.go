package utils

import (
	"testing"
)

func TestModifyFileExt(t *testing.T) {

	type args struct {
		filename string
		newExt   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{filename: "/path/to/1.mp4", newExt: "mp3"}, want: "/path/to/1.mp3"},
		{args: args{filename: "./path/to/1.mp4", newExt: "mp3"}, want: "./path/to/1.mp3"},
		{args: args{filename: "/path/to/1", newExt: "mp3"}, want: "/path/to/1.mp3"},
		{args: args{filename: "/path/to/1..mp4", newExt: "mp3"}, want: "/path/to/1..mp3"},
		{args: args{filename: "/path/to/1.mp3.mp4", newExt: "mp3"}, want: "/path/to/1.mp3.mp3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ModifyFileExt(tt.args.filename, tt.args.newExt); got != tt.want {
				t.Errorf("ModifyFileExt() = %v, want %v", got, tt.want)
			}
		})
	}
}
