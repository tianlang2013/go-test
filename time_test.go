package study

import (
	"time"
	"fmt"
	"testing"
)

func TestTime(t *testing.T) {
	//NewTiker指定的是多久产生一次
	var ticker *time.Ticker = time.NewTicker(2 * time.Second)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Second * 5) //阻塞，则执行次数为sleep的休眠时间/ticker的时间
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
