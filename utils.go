package go_concurrency_models

import (
	"os"
	"fmt"
)

func PanicOnErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func OpenFileOrPanic(filePath string) *os.File {
	handle, err := os.Open(filePath)
	PanicOnErr(err)
	return handle
}
