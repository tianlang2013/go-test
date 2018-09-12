package study

import (
	"testing"
	"math/big"
	"fmt"
	"reflect"
)

func TestBig(t *testing.T)  {
	big1 := new(big.Int).SetUint64(uint64(1000))
	fmt.Println("big1 is: ", big1)

	big2 := big1.Uint64()
	fmt.Println("big2 is: ", big2)

	//big3 := big1.Int64()
	fmt.Println("big3 is: ", reflect.TypeOf(big2))

}


func TestByte(t *testing.T){
	big1 ,_:= new(big.Int).SetString("1000", 10)
	fmt.Println("big1 is: ", big1)

	big2 := big1.Bytes()
	fmt.Println("big2 is: ", big2)

}
