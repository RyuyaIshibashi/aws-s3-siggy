package main

import (
	"aws-s3-siggy/cmd"
)

var (
	version  string
	revision string
)

func main() {
	cmd.Execute(version, revision)
}
