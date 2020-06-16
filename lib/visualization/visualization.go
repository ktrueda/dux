package visualization

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/dustin/go-humanize"
)

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

func showLine(label string, val int64, length int, unit string, labelFomat string) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf(labelFomat, label))
	for i := 0; i < length; i++ {
		buffer.WriteString("▇")
	}
	buffer.WriteString(fmt.Sprintf(" %s %s", humanize.Comma(val), unit))
	return buffer.String()
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
