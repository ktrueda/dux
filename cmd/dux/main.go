package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ktrueda/dux/pkg/util"

	"github.com/dustin/go-humanize"
	"github.com/ttacon/chalk"
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

	fi, errStat := os.Stat(root)
	if errStat != nil {
		panic(errStat)
	}
	if !fi.Mode().IsDir() {
		fmt.Fprintln(os.Stderr, root, " is not directory.")
		os.Exit(1)
	}

	var suffixSizeMap, directorySizeMap, topLargeFiles = util.Inspect(root)

	sectionStyle := chalk.Red.NewStyle().
		WithBackground(chalk.Black).
		WithTextStyle(chalk.Bold).
		WithTextStyle(chalk.Underline).
		Style

	fmt.Println(sectionStyle("🐘 File Size Group By suffix"))
	util.Show(suffixSizeMap)
	fmt.Println("")
	fmt.Println(sectionStyle("🦒 File Size Group By directory"))
	util.Show(directorySizeMap)

	fmt.Println("")
	fmt.Println(sectionStyle("🦛 Top Large size file"))
	for i := 0; i < len(topLargeFiles); i++ {
		f := topLargeFiles[i]
		if f.Value > 0 {
			fmt.Println(fmt.Sprintf("%d %s %s B", i, f.Key, humanize.Comma(f.Value)))
		}
	}
}
