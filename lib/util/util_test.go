package util

import (
	"fmt"
	"github.com/google/uuid"
	"os"
	"testing"
)

func TestChildDir(t *testing.T) {
	type args struct {
		path string
		root string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"/dir1/dir2/file", "/dir1"}, "/dir1/dir2"},
		{"test2", args{"/dir1/file", "/dir1"}, "/dir1"},
		{"test3", args{"/dir1/dir2/file", "/dir1/dir2"}, "/dir1/dir2"},
		{"test4", args{"/dir1/dir2/dir3/file", "/dir1/dir2"}, "/dir1/dir2/dir3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utilChildDir(tt.args.path, tt.args.root); got != tt.want {
				t.Errorf("ChildDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileSize(t *testing.T) {
	uuidObj, _ := uuid.NewUUID()
	var tempFile = fmt.Sprintf("/tmp/%s", uuidObj.String())

	f, _ := os.Create(tempFile)
	f.WriteString("hello")
	f.Close()

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{"exist", args{tempFile}, 5, false},
		{"not found", args{"/tmp/not_cound"}, -1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := util.FileSize(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FileSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
