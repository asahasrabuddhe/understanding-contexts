// START MAIN OMIT
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(100*time.Microsecond, func() {
		cancel()
	})

	countedUntil := DoTask(ctx)
	fmt.Println("counted until", countedUntil)
}

// END MAIN OMIT
// START FUNC OMIT
func DoTask(ctx context.Context) int {
	var num int
	var finished bool

	fmt.Println("working...")

	for !finished {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			finished = true
		default:
			num++
		}
	}

	return num
}

// END FUNC OMIT
