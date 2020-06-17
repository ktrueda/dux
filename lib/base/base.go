package base

import (
	"fmt"
	"os"
)

type File struct {
	Path string
	Size int64
}
type FileList []File

func (t FileList) Len() int {
	return len(t)
}
func (t FileList) Less(i, j int) bool {
	return t[i].Size < t[j].Size
}
func (t FileList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (h *FileList) Push(x interface{}) {
	*h = append(*h, x.(File))
}

func (h *FileList) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Stderr(verbose bool, text ...interface{}) {
	if verbose {
		fmt.Fprintln(os.Stderr, text...)
	}
}
