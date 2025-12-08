package inputreader

import (
	"fmt"
	"os"
)

func Readfile(InputFile string) (string, error) {
	InputContent, err := os.ReadFile(InputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read file %q:%w", InputFile, err)
	}
	return string(InputContent), nil
}
