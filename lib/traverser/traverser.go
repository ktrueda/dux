package traverser

import (
	"os"
	"path/filepath"

	"github.com/ktrueda/dux/lib/base"
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
func Inspect(root string, verbose bool) (map[string]int64, map[string]int64, [10]base.File) {
	var suffixSizeMap map[string]int64 = map[string]int64{}
	var directorySizeMap map[string]int64 = map[string]int64{}
	largeFiles := [10]base.File{}
	for i := 0; i < 10; i++ {
		largeFiles[i] = base.File{Path: "", Size: -1}
	}

	errWalk := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		fi, errStat := os.Stat(path)
		if errStat != nil {
			base.Stderr(verbose, "Skip", path, " (cannot get stat)")
			return nil
		}
		if fi.Mode().IsDir() {
			return nil
		}
		var suffix string = filepath.Ext(path)
		var dir string = ChildDir(path, root)
		size, errSize := FileSize(path)
		if errSize != nil {
			base.Stderr(verbose, "Skip", path, " (cannot get size)")
			return nil
		}
		suffixSizeMap[suffix] += size
		directorySizeMap[dir] += size
		updateLargestArray(&largeFiles, base.File{Path: path, Size: size})

		return nil
	})
	if errWalk != nil {
		panic(errWalk)
	}

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

func updateLargestArray(minList *[10]base.File, newValue base.File) {
	/**
	update array
	{
		base.File("/path/to/file9", 150),
		...
		base.File("/path/to/file3", 125),
		base.File("/path/to/file2", 124),
		base.File("/path/to/file1", 123),
		base.File("/path/to/file0", 122),
	}
	*/
	var size int = len(*minList)
	if (*minList)[size-1].Size < newValue.Size {
		(*minList)[size-1] = newValue
	} else {
		return
	}

	for i := size - 1; i > 0; i-- {
		if (*minList)[i].Size > (*minList)[i-1].Size {
			(*minList)[i], (*minList)[i-1] = (*minList)[i-1], (*minList)[i]
		} else {
			return
		}
	}
}
