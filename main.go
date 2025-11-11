package main

import (
	"os"

	"platform.zone01.gr/git/vxanthio/go-reloaded/inputreader"
	"platform.zone01.gr/git/vxanthio/go-reloaded/tokenizer"
)

func main() {
	InputFile := os.Args[1]
	OutputFile := os.Args[2]
	Content := inputreader.Readfile(InputFile)
	os.WriteFile(OutputFile, []byte(Content), 0644)
	tokens := tokenizer.Tokenize(Content)
}
