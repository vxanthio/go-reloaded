package main

import (
	"fmt"
	"os"

	"platform.zone01.gr/git/vxanthio/go-reloaded/internal/inputreader"
	"platform.zone01.gr/git/vxanthio/go-reloaded/internal/tokenizer"
)

func main() {
	InputFile := os.Args[1]
	OutputFile := os.Args[2]
	Content := inputreader.Readfile(InputFile)
	tokens := tokenizer.Tokenize(Content)
	fmt.Println(tokens)
	os.WriteFile(OutputFile, []byte(Content), 0644)
}
