package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dustin/go-humanize"
	"github.com/ktrueda/dux/lib/traverser"
	"github.com/ktrueda/dux/lib/visualization"
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

	var suffixSizeMap, directorySizeMap, topLargeFiles = traverser.Inspect(root)

	sectionStyle := chalk.Red.NewStyle().
		WithBackground(chalk.Black).
		WithTextStyle(chalk.Bold).
		WithTextStyle(chalk.Underline).
		Style

	fmt.Println(sectionStyle("🐘 File Size Group By suffix"))
	fmt.Println(visualization.GraphString(suffixSizeMap))
	fmt.Println("")
	fmt.Println(sectionStyle("🦒 File Size Group By directory"))
	fmt.Println(visualization.GraphString(directorySizeMap))

	fmt.Println("")
	fmt.Println(sectionStyle("🦛 Top Large size file"))
	for i := 0; i < len(topLargeFiles); i++ {
		f := topLargeFiles[i]
		if f.Size > 0 {
			fmt.Println(fmt.Sprintf("%d %s %s B", i, f.Path, humanize.Comma(f.Size)))
		}
	}
}
