package inputreader

import (
	"fmt"
	"os"
)

func Readfile(InputFile string) string {
	InputContent, err := os.ReadFile(InputFile)
	if err != nil {

		fmt.Println("Σφαλμα δεν μπορεσα να ανοιξω το αρχειο ελεγχου")
		return ""
	}
	return string(InputContent)
}
