// START MAIN OMIT
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()
	if deadline, ok := ctx.Deadline(); ok {
		fmt.Println("program will terminate at", deadline)
	}
	countedUntil := DoTask(ctx)
	fmt.Println("counted until", countedUntil)
}

// END MAIN OMIT
// START FUNC OMIT
func DoTask(ctx context.Context) int {
	var num int

	fmt.Println("working...")

	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return num
		default:
			num++
		}
	}
}

// END FUNC OMIT
