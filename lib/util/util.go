package util

import (
	"bytes"
	"container/heap"
	"fmt"
	"github.com/dustin/go-humanize"
	"os"
	"path/filepath"
	"sort"
)

/*
return file size
*/
func FileSize(path string) (int64, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return -1, err
	}
	return fi.Size(), nil
}

/*
return one-parent directory path
e.g.
/dir1/dir2/file -> /dir1/dir2
*/
func ChildDir(path string, root string) string {
	return recursiveChildDir(path, root, false)
}

/*
Inspect given root path
return something about the root path
*/
func Inspect(root string) (map[string]int64, map[string]int64, []File) {
	var suffixSizeMap map[string]int64 = map[string]int64{}
	var directorySizeMap map[string]int64 = map[string]int64{}

	// heap to store large files
	largeFileHeap := &FileList{}
	heap.Init(largeFileHeap)

	errWalk := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		fi, errStat := os.Stat(path)
		if errStat != nil {
			fmt.Println(fmt.Printf("Skip %s (cannot get stat)", path))
			return nil
		}
		if fi.Mode().IsDir() {
			return nil
		}
		var suffix string = filepath.Ext(path)
		var dir string = ChildDir(path, root)
		size, errSize := FileSize(path)
		if errSize != nil {
			fmt.Println(fmt.Printf("Skip %s (cannot get size)", path))
			return nil
		}
		suffixSizeMap[suffix] += size
		directorySizeMap[dir] += size
		heap.Push(largeFileHeap, File{path, size * -1})

		if largeFileHeap.Len() >= 11 {
			heap.Remove(largeFileHeap, 10)
		}

		return nil
	})
	if errWalk != nil {
		panic(errWalk)
	}

	var arr = make([]File, largeFileHeap.Len())
	for i := 0; i < len(arr); i++ {
		arr[i] = heap.Pop(largeFileHeap).(File)
	}

	return suffixSizeMap, directorySizeMap, arr
}

func recursiveChildDir(path string, root string, rec bool) string {
	var parent = filepath.Dir(path)
	// fmt.Println(fmt.Sprintf("root: %s path: %s parent: %s", root, path, parent))
	if parent == root {
		if rec {
			return path
		} else {
			return parent
		}
	} else {
		return recursiveChildDir(parent, root, true)
	}
}

type File struct {
	Key   string
	Value int64
}
type FileList []File

func (t FileList) Len() int {
	return len(t)
}
func (t FileList) Less(i, j int) bool {
	return t[i].Value < t[j].Value
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

func normalize(data map[string]int64) map[string]int {
	var maxv int64
	for _, v := range data {
		if maxv < v {
			maxv = v
		}
	}
	var normalized map[string]int = map[string]int{}
	for k, v := range data {
		normalized[k] = int(50 * v / maxv)
	}
	return normalized
}

func Show(data map[string]int64) {
	normalizedData := normalize(data)

	// sort by value
	sorter := make(FileList, len(data))
	index := 0
	for k, v := range data {
		sorter[index] = File{k, v}
		index++
	}
	sort.Sort(sort.Reverse(sorter))

	// label max length
	var maxLabelLength = 10
	for k, _ := range data {
		if maxLabelLength < len(k) {
			maxLabelLength = len(k)
		}
	}
	var labelFormat = "%" + fmt.Sprintf("%d", maxLabelLength) + "s : "

	// show each line
	for _, d := range sorter {
		showLine(d.Key, d.Value, normalizedData[d.Key], "B", labelFormat)
	}
}

func showLine(label string, val int64, length int, unit string, labelFomat string) {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf(labelFomat, label))
	for i := 0; i < length; i++ {
		buffer.WriteString("▇")
	}
	buffer.WriteString(fmt.Sprintf(" %s %s", humanize.Comma(val), unit))
	fmt.Println(buffer.String())
}
