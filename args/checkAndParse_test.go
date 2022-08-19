package args

import (
	"testing"
)

func Test_checkOpts(t *testing.T) {
	type args struct {
		opt *Options
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{args: args{opt: &Options{Out: "/tmp/mp3"}}, wantErr: false},
		{args: args{opt: &Options{Out: "/tmp/mp3/1"}}, wantErr: false},
		{args: args{opt: &Options{Out: "/tmp/mp3/1/"}}, wantErr: true},
	}
	for _, tt := range tests {
		t.Skipf(tt.name, func(t *testing.T) {
			if err := checkAndAdjustOpts(tt.args.opt); (err != nil) != tt.wantErr {
				t.Errorf("checkAndAdjustOpts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_parseSearchSource(t *testing.T) {
	type args struct {
		opt *Options
	}
	tests := []struct {
		name        string
		args        args
		wantSources []string
	}{
		{
			name: "all field not specified, should use default: all",
			args: args{opt: &Options{Search: Search{Sources: "", ExcludeSources: ""}}},
			wantSources: []string{
				"bilibili", "douyin", "kugou", "kuwo", "migu", "netease", "qq", "youtube", "qmkg",
			},
		},
		{
			name: "sources contain unknown value",
			args: args{opt: &Options{Search: Search{Sources: "kuwo,fakeSource", ExcludeSources: ""}}},
			wantSources: []string{
				"kuwo",
			},
		},
		{
			name: "sources specified but not excludeSources",
			args: args{opt: &Options{Search: Search{Sources: "kuwo,qq", ExcludeSources: ""}}},
			wantSources: []string{
				"kuwo", "qq",
			},
		},
		{
			name: "sources specified with space but not excludeSources",
			args: args{opt: &Options{Search: Search{Sources: "kuwo, qq", ExcludeSources: ""}}},
			wantSources: []string{
				"kuwo", "qq",
			},
		},
		{
			name: "excludeSources specified but not sources",
			args: args{opt: &Options{Search: Search{Sources: "", ExcludeSources: "kuwo,qq"}}},
			wantSources: []string{
				"bilibili", "douyin", "kugou", "migu", "netease", "youtube", "qmkg",
			},
		},
		{
			name: "excludeSources with space specified but not sources",
			args: args{opt: &Options{Search: Search{Sources: "", ExcludeSources: "kuwo , qq"}}},
			wantSources: []string{
				"bilibili", "douyin", "kugou", "migu", "netease", "youtube", "qmkg",
			},
		},
		{
			name: "both are specified",
			args: args{opt: &Options{Search: Search{Sources: "kuwo,qq", ExcludeSources: "migu,netease"}}},
			wantSources: []string{
				"kuwo", "qq",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseSearchSource(tt.args.opt)
			if !isSameWithTheTwoSlice(tt.args.opt.Search.SourcesWillBeSearch, tt.wantSources) {
				t.Errorf("want: %v, but got %v", tt.wantSources, tt.args.opt.Search.SourcesWillBeSearch)
			}
		})
	}
}

func isSameWithTheTwoSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
