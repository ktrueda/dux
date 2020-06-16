package traverser

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/ktrueda/dux/lib/base"
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
			if got := ChildDir(tt.args.path, tt.args.root); got != tt.want {
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
			got, err := FileSize(tt.args.path)
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

func Test_updateLargetArray(t *testing.T) {
	type args struct {
		minList  *[10]base.File
		newValue base.File
	}
	tests := []struct {
		name string
		args args
		want [10]base.File
	}{
		{
			"add 1 remove1",
			args{
				&[10]base.File{
					base.File{"/tmp/file9", 90},
					base.File{"/tmp/file8", 80},
					base.File{"/tmp/file7", 70},
					base.File{"/tmp/file6", 60},
					base.File{"/tmp/file5", 50},
					base.File{"/tmp/file4", 40},
					base.File{"/tmp/file3", 30},
					base.File{"/tmp/file2", 20},
					base.File{"/tmp/file1", 10},
					base.File{"/tmp/file0", 0},
				},
				base.File{"/tmp/file_new", 123},
			},
			[10]base.File{
				base.File{"/tmp/file_new", 123},
				base.File{"/tmp/file9", 90},
				base.File{"/tmp/file8", 80},
				base.File{"/tmp/file7", 70},
				base.File{"/tmp/file6", 60},
				base.File{"/tmp/file5", 50},
				base.File{"/tmp/file4", 40},
				base.File{"/tmp/file3", 30},
				base.File{"/tmp/file2", 20},
				base.File{"/tmp/file1", 10},
			},
		},
		{
			"stay",
			args{
				&[10]base.File{
					base.File{"/tmp/file9", 90},
					base.File{"/tmp/file8", 80},
					base.File{"/tmp/file7", 70},
					base.File{"/tmp/file6", 60},
					base.File{"/tmp/file5", 50},
					base.File{"/tmp/file4", 40},
					base.File{"/tmp/file3", 30},
					base.File{"/tmp/file2", 20},
					base.File{"/tmp/file1", 10},
					base.File{"/tmp/file0", 5},
				},
				base.File{"/tmp/file_new", 2},
			},
			[10]base.File{
				base.File{"/tmp/file9", 90},
				base.File{"/tmp/file8", 80},
				base.File{"/tmp/file7", 70},
				base.File{"/tmp/file6", 60},
				base.File{"/tmp/file5", 50},
				base.File{"/tmp/file4", 40},
				base.File{"/tmp/file3", 30},
				base.File{"/tmp/file2", 20},
				base.File{"/tmp/file1", 10},
				base.File{"/tmp/file0", 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateLargestArray(tt.args.minList, tt.args.newValue)
			if *tt.args.minList != tt.want {
				t.Errorf("actual = %v, want %v", *tt.args.minList, tt.want)
			}
		})
	}
}
