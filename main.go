package main

import (
	"./dux"
	"flag"
	"fmt"
	"github.com/dustin/go-humanize"
	"os"
	"path/filepath"
)

func main() {

	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	root, err1 := filepath.Abs(flag.Args()[0])
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(fmt.Sprintf("Target Directory: %s", root))

	var suffixSizeMap, directorySizeMap, topLargeFiles = dux.Inspect(root)

	fmt.Println("File Size Group By suffix")
	dux.Show(suffixSizeMap)
	fmt.Println("")
	fmt.Println("File Size Group By directory")
	dux.Show(directorySizeMap)

	fmt.Println("")
	fmt.Println("Top Large size file")
	for i := 0; i < len(topLargeFiles); i++ {
		f := topLargeFiles[i]
		fmt.Println(fmt.Sprintf("%s %s B", f.Key, humanize.Comma(f.Value*-1)))
	}
}
