package study

import (
	"testing"
	"fmt"
	"bufio"
	"os"
)

func TestB1(t *testing.T){
	var b []byte
	b = []byte("abc")
	b = append(b , 10)
	fmt.Println(b)


}

func TestBufferIo(t *testing.T){
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
