package utils

import (
	"testing"
)

func TestConvertString2Int(t *testing.T) {
	type args struct {
		numberStr string
		fallback  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				numberStr: "123",
				fallback:  1,
			},
			want: 123,
		},
		{
			args: args{
				numberStr: "123.3",
				fallback:  1,
			},
			want: 123,
		},
		{
			args: args{
				numberStr: "ff",
				fallback:  1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertString2Int(tt.args.numberStr, tt.args.fallback); got != tt.want {
				t.Errorf("ConvertString2Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
