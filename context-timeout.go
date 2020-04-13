package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Microsecond)
	defer cancel()
	if deadline, ok := ctx.Deadline(); ok {
		fmt.Println("program will terminate at", deadline)
	}
	countedUntil := DoTask(ctx)
	fmt.Println("counted until", countedUntil)
}

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
