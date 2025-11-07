package main

import (
	"os"

	"platform.zone01.gr/git/vxanthio/go-reloaded/inputreader"
)

func main() {
	InputFile := os.Args[1]
	OutputFile := os.Args[2]
	Content := inputreader.Readfile(InputFile)
	os.WriteFile(OutputFile, []byte(Content), 0644)
}
