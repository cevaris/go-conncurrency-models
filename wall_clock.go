package go_concurrency_models

import (
	"fmt"
	"time"
)


type WallClock struct {
	StartTime int64
	EndTime int64
	DurationTime int64
}

func NewWallClock() *WallClock {
	return &WallClock{}
}

// Start clock
func (wc *WallClock) StartClock() {
	wc.StartTime = time.Now().Unix()
}

// Stop running clock
func (wc *WallClock) StopClock() {
	wc.EndTime = time.Now().Unix()
	wc.DurationTime = wc.EndTime - wc.StartTime
	fmt.Println("Total Time (sec):", wc.DurationTime)
}
