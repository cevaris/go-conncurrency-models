package wiki

import (
	"os"
	"fmt"
)
func OpenFileHandler(filePath string) (*os.File, error) {
	handle, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file.", err)
		return nil, err
	}
	return handle, nil
}
