package main

import (
	"flag"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/pkg/profile"
	"local.packages/util"
	"os"
	"path/filepath"
)

func main() {
	defer profile.Start(profile.ProfilePath(".")).Stop()

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

	var suffixSizeMap, directorySizeMap, topLargeFiles = util.Inspect(root)

	fmt.Println("File Size Group By suffix")
	util.Show(suffixSizeMap)
	fmt.Println("")
	fmt.Println("File Size Group By directory")
	util.Show(directorySizeMap)

	fmt.Println("")
	fmt.Println("Top Large size file")
	for i := 0; i < len(topLargeFiles); i++ {
		f := topLargeFiles[i]
		fmt.Println(fmt.Sprintf("%d %s %s B", i, f.Key, humanize.Comma(f.Value)))
	}
}
