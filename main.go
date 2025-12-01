package main

import (
	"fmt"
	"os"

	"platform.zone01.gr/git/vxanthio/go-reloaded/internal/inputreader"
	"platform.zone01.gr/git/vxanthio/go-reloaded/internal/ruleprocessor"
	"platform.zone01.gr/git/vxanthio/go-reloaded/internal/tokenizer"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage:go run main.go <input_file> <output_file>")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	text, err := inputreader.Readfile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	tokens := tokenizer.Tokenize(text)
	processed := ruleprocessor.ProcessTokens(tokens)
	finalText := formatter.BuildOutput(processed)
err=os.WriteFile(outputFile[]byte(finalText),0644)
if err!=nil {
fmt.Println("Error writing output:",err)
return 
}
}
