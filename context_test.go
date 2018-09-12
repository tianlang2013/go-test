package study

import (
	"testing"
	"time"
	"fmt"
	"context"
)

func TestDeadLine(t *testing.T) {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				fmt.Println("overslept")
			case <-ctx.Done():
				fmt.Println("1" ,ctx.Err())
				return
			}
		}
	}()

	go func() {
		for {
			select {

			case <-ctx.Done():
				fmt.Println("2",ctx.Err())
				return
			}
		}
	}()
	for {
		select {

		case <-time.After(1 * time.Second):
			fmt.Println("3",ctx.Err())
		}
	}
}
