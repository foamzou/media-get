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

func TestSimilarText(t *testing.T) {
	type args struct {
		str1   string
		str2   string
		target string
	}
	tests := []struct {
		name        string
		args        args
		wantCompare string
	}{
		{
			name: "test [=] chinese. same",
			args: args{
				str1:   "一样的",
				str2:   "一样的",
				target: "没任何相同",
			},
			wantCompare: "=",
		},
		{
			name: "test [=] chinese",
			args: args{
				str1:   "都不一样",
				str2:   "是的发发",
				target: "天空",
			},
			wantCompare: "=",
		},
		{
			name: "test [=] chinese and english",
			args: args{
				str1:   "Initial J",
				str2:   "中国新歌声第十五期",
				target: "范特西",
			},
			wantCompare: "=",
		},
		{
			name: "test [=] empty vs not empty",
			args: args{
				str1:   "",
				str2:   "方法",
				target: "范特西",
			},
			wantCompare: "=",
		},
		{
			name: "test [>]。 精准匹配",
			args: args{
				str1:   "范特西",
				str2:   "无所谓什么名字",
				target: "范特西",
			},
			wantCompare: ">",
		},
		{
			name: "test [>]。 精准名字 vs 后缀",
			args: args{
				str1:   "范特西",
				str2:   "范特西1",
				target: "范特西",
			},
			wantCompare: ">",
		},
		{
			name: "test [>]。 精准名字 vs 前缀",
			args: args{
				str1:   "范特西",
				str2:   "2范特西",
				target: "范特西",
			},
			wantCompare: ">",
		},
		{
			name: "test [>]。 前缀 vs 后缀",
			args: args{
				str1:   "不能说的秘密是什么",
				str2:   "哈哈哈不能说的秘密",
				target: "不能说的秘密",
			},
			wantCompare: ">",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str1Ret := SimilarText(tt.args.target, tt.args.str1)
			str2Ret := SimilarText(tt.args.target, tt.args.str2)
			compare := ""
			if str1Ret == str2Ret {
				compare = "="
			} else if str1Ret > str2Ret {
				compare = ">"
			} else {
				compare = "<"
			}

			if compare != tt.wantCompare {
				t.Errorf("for target: %v, want %v %v %v, but got %v %v %v", tt.args.target, tt.args.str1, tt.wantCompare, tt.args.str2, str1Ret, compare, str2Ret)
			}
		})
	}
}
