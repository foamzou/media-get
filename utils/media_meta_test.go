package utils

import (
	"testing"
)

func TestDurationStr2Second(t *testing.T) {
	type args struct {
		durationStr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{durationStr: "00:12"},
			want: 12,
		},
		{
			args: args{durationStr: "01:12"},
			want: 72,
		},
		{
			args: args{durationStr: "00:00:12"},
			want: 12,
		},
		{
			args: args{durationStr: "00:01:12"},
			want: 72,
		},
		{
			args: args{durationStr: "01:01:12"},
			want: 3672,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DurationStr2Second(tt.args.durationStr); got != tt.want {
				t.Errorf("DurationStr2Second() = %v, want %v", got, tt.want)
			}
		})
	}
}
