// +build windows

package main

import (
	"errors"
	"testing"
)

func Test_execute(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name   string
		zapper func(string) error
		args   args
		want   int
	}{
		{
			name:   "no folder specified",
			zapper: func(string) error { return nil },
			args: args{
				args: []string{"exeName"},
			},
			want: 1,
		},
		{
			name:   "folder specified",
			zapper: func(string) error { return nil },
			args: args{
				args: []string{"exeName", "-folder", "testfolder"},
			},
			want: 0,
		},
		{
			name:   "invalid flag",
			zapper: func(string) error { return nil },
			args: args{
				args: []string{"exeName", "-invalid", "testfolder"},
			},
			want: 1,
		},
		{
			name:   "zap failure",
			zapper: func(string) error { return errors.New("test error") },
			args: args{
				args: []string{"exeName", "-folder", "testfolder"},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zapper = tt.zapper
			if got := execute(tt.args.args); got != tt.want {
				t.Errorf("execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
