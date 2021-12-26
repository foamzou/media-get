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
		t.Run(tt.name, func(t *testing.T) {
			if err := checkAndAdjustOpts(tt.args.opt); (err != nil) != tt.wantErr {
				t.Errorf("checkAndAdjustOpts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
