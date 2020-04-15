// START MAIN OMIT
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Microsecond)
	defer cancel()

	countedUntil := DoTask(ctx)
	fmt.Println("counted until", countedUntil)
}

// END MAIN OMIT
// START FUNC OMIT
func DoTask(ctx context.Context) int {
	var num int
	var finished bool
	if deadline, ok := ctx.Deadline(); ok { // HL
		// this can be an excellent opportunity to check if we have // HL
		// enough time on hand to even initiate the operation // HL
		fmt.Println("program will terminate at", deadline) // HL
	} // HL
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
