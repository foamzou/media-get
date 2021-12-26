package processor

import (
	"testing"

	args2 "github.com/foamzou/audio-get/args"
)

func Test_getOutputPaths(t *testing.T) {
	type args struct {
		opts      *args2.Options
		fileTitle string
	}
	tests := []struct {
		name              string
		args              args
		wantTempFilePath  string
		wantTargetOutPath string
	}{
		{
			args: args{
				opts:      &args2.Options{Out: "/path/to/1.mp3", IsDir: false},
				fileTitle: "dream",
			},
			wantTempFilePath: "/path/to/dream", wantTargetOutPath: "/path/to/1.mp3",
		},
		{
			args: args{
				opts:      &args2.Options{Out: "/path/to/1.audio", IsDir: false},
				fileTitle: "dream",
			},
			wantTempFilePath: "/path/to/dream", wantTargetOutPath: "/path/to/1.audio",
		},
		{
			args: args{
				opts:      &args2.Options{Out: "/path/to", IsDir: true},
				fileTitle: "dream",
			},
			wantTempFilePath: "/path/to/dream", wantTargetOutPath: "/path/to/dream.mp3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Processor{Opts: tt.args.opts}
			gotTempFilePath, gotTargetOutPath := p.getOutputPaths(tt.args.fileTitle)
			if gotTempFilePath != tt.wantTempFilePath {
				t.Errorf("getOutputPaths() gotTempFilePath = %v, want %v", gotTempFilePath, tt.wantTempFilePath)
			}
			if gotTargetOutPath != tt.wantTargetOutPath {
				t.Errorf("getOutputPaths() gotTargetOutPath = %v, want %v", gotTargetOutPath, tt.wantTargetOutPath)
			}
		})
	}
}
