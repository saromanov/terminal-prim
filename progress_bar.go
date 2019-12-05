package prim

import (
	"time"

	"github.com/cheggaaa/pb"
)

// ProgressBar creates a basic progress bar
func ProgressBar(count int) {
	bar := pb.StartNew(count)
	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.Finish()
}
