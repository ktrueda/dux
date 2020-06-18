module github.com/ktrueda/dux

go 1.14

replace github.com/ktrueda/dux/lib/traverser => ./lib/traverser

replace github.com/ktrueda/dux/lib/visualization => ./lib/visualization

replace github.com/ktrueda/dux/lib/base => ./lib/base

require (
	github.com/dustin/go-humanize v1.0.0
	github.com/ktrueda/dux/lib/traverser v0.0.0-00010101000000-000000000000
	github.com/ktrueda/dux/lib/visualization v0.0.0-00010101000000-000000000000
	github.com/pkg/profile v1.5.0
	github.com/ttacon/chalk v0.0.0-20160626202418-22c06c80ed31
	github.com/urfave/cli/v2 v2.2.0
)
