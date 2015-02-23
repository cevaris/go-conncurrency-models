package go_concurrency_models

import (
	"fmt"
	"os"
	"time"
)

var startTime int64

func StartClock() {
	startTime = time.Now().Unix()
}

func StopClock() {
	endTime := time.Now().Unix()
	duration := endTime - startTime
	fmt.Println(duration)
}

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
