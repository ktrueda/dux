package main

import (
	"fmt"
	"os"

	"github.com/dustin/go-humanize"
	"github.com/ktrueda/dux/lib/traverser"
	"github.com/ktrueda/dux/lib/visualization"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli/v2"
)

func main() {
	var verbose bool
	app := &cli.App{
		Name:  "dux",
		Usage: "du eXtended",
		Action: func(c *cli.Context) error {
			var targetPath string = c.Args().First()
			fmt.Println(fmt.Sprintf("Target Directory: %s", targetPath))

			fi, errStat := os.Stat(targetPath)
			if errStat != nil {
				panic(errStat)
			}
			if !fi.Mode().IsDir() {
				fmt.Fprintln(os.Stderr, targetPath, " is not directory.")
				os.Exit(1)
			}
			execute(targetPath, verbose)
			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "verbose",
				Usage:       "Make the operation more talkative",
				Destination: &verbose,
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func execute(root string, verbose bool) {

	var suffixSizeMap, directorySizeMap, topLargeFiles = traverser.Inspect(root, verbose)

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
