package utils

import "testing"

func TestExecCmd(t *testing.T) {
	type args struct {
		cmd  string
		arg1 string
		arg2 string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Test echo",
			args:    args{cmd: "echo", arg1: "hello", arg2: "foam"},
			want:    "hello foam\n",
			wantErr: false,
		},
		{
			name:    "Test error",
			args:    args{cmd: "error"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExecCmd(tt.args.cmd, tt.args.arg1, tt.args.arg2)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecCmd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExecCmd() got = %v, want %v", got, tt.want)
			}
		})
	}
}
