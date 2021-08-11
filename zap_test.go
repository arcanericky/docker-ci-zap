// +build windows

package dockercizap

import (
	"errors"
	"os"
	"testing"

	"github.com/Microsoft/hcsshim"
)

func Test_folderExists(t *testing.T) {
	type args struct {
		folder string
	}
	tests := []struct {
		name   string
		create func()
		remove func()
		args   args
		want   bool
	}{
		{
			name: "folder exists",
			create: func() {
				os.Mkdir("testfolder", os.FileMode(0755))
			},
			remove: func() {
				os.Remove("testfolder")
			},
			args: args{
				folder: "testfolder",
			},
			want: true,
		},
		{
			name:   "folder does not exist",
			create: func() {},
			remove: func() {},
			args: args{
				folder: "testfolder",
			},
			want: false,
		},
		{
			name: "folder a file (not a dir)",
			create: func() {
				f, _ := os.Create("testfolder")
				f.Close()
			},
			remove: func() {
				os.Remove("testfolder")
			},
			args: args{
				folder: "testfolder",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.create()
			if got := folderExists(tt.args.folder); got != tt.want {
				t.Errorf("folderExists() = %v, want %v", got, tt.want)
			}
			tt.remove()
		})
	}
}

func Test_destroyLayer(t *testing.T) {
	type args struct {
		destroyer func(hcsshim.DriverInfo, string) error
		folder    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "destroy success",
			args: args{
				destroyer: func(hcsshim.DriverInfo, string) error {
					return nil
				},
				folder: "testfolder",
			},
			wantErr: false,
		},
		{
			name: "destroy fail",
			args: args{
				destroyer: func(hcsshim.DriverInfo, string) error {
					return errors.New("test error")
				},
				folder: "testfolder",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			destroyer = tt.args.destroyer
			if err := destroyLayer(tt.args.folder); (err != nil) != tt.wantErr {
				t.Errorf("destroyLayer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Zap(t *testing.T) {
	type args struct {
		destroyer     func(hcsshim.DriverInfo, string) error
		folderChecker func(string) bool
		folder        string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				destroyer: func(hcsshim.DriverInfo, string) error {
					return nil
				},
				folderChecker: func(string) bool {
					return true
				},
			},
			wantErr: false,
		},
		{
			name: "folder does not exist",
			args: args{
				destroyer: func(hcsshim.DriverInfo, string) error {
					return nil
				},
				folderChecker: func(string) bool {
					return false
				},
			},
			wantErr: true,
		},
		{
			name: "destroy fails",
			args: args{
				destroyer: func(hcsshim.DriverInfo, string) error {
					return errors.New("test error")
				},
				folderChecker: func(string) bool {
					return true
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			destroyer = tt.args.destroyer
			folderChecker = tt.args.folderChecker
			if err := Zap(tt.args.folder); (err != nil) != tt.wantErr {
				t.Errorf("zap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
