package util

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/theckman/yacspin"
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
func Inspect(root string) (map[string]int64, map[string]int64, [10]File) {
	var suffixSizeMap map[string]int64 = map[string]int64{}
	var directorySizeMap map[string]int64 = map[string]int64{}
	largeFiles := [10]File{}

	cfg := yacspin.Config{
		Frequency:  100 * time.Millisecond,
		CharSet:    yacspin.CharSets[59],
		Suffix:     "Checking ",
		StopColors: []string{"fgGreen"},
	}

	spinner, _ := yacspin.New(cfg)

	spinner.Start()
	errWalk := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		spinner.Message(path)
		fi, errStat := os.Stat(path)
		if errStat != nil {
			fmt.Fprintln(os.Stderr, "Skip", path, " (cannot get stat)")
			return nil
		}
		if fi.Mode().IsDir() {
			return nil
		}
		var suffix string = filepath.Ext(path)
		var dir string = ChildDir(path, root)
		size, errSize := FileSize(path)
		if errSize != nil {
			fmt.Fprintln(os.Stderr, "Skip", path, " (cannot get size)")
			return nil
		}
		suffixSizeMap[suffix] += size
		directorySizeMap[dir] += size
		updateMinList(&largeFiles, File{path, size})

		return nil
	})
	if errWalk != nil {
		panic(errWalk)
	}
	spinner.Stop()

	return suffixSizeMap, directorySizeMap, largeFiles
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
		normalized[k] = int(50 * v / (maxv + 1))
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

func updateMinList(minList *[10]File, newValue File) {
	size := len(*minList)
	if (*minList)[size-1].Value < newValue.Value {
		(*minList)[size-1] = newValue
	} else {
		return
	}

	for i := 0; i < size-1; i++ {
		if (*minList)[9-i].Value > (*minList)[8-i].Value {
			(*minList)[9-i], (*minList)[8-i] = (*minList)[8-i], (*minList)[9-i]
		} else {
			return
		}
	}
}
