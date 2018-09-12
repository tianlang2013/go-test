package study

import (
	"testing"
	"time"
	"fmt"
)

/**
	select 默认是阻塞的，直到channel有数据返回
	所以此时for循环并没有一直运行，只是在等待状态

 */
func TestSelect(t *testing.T)  {

	for {
		timer :=time.NewTimer(time.Second * 5)
		select {
		case <- timer.C  :
			fmt.Println("ss")
		default:
			fmt.Printf("no one was ready to communicate\n")
		}
		fmt.Println("f")
	}

}
