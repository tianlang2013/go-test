package study

import (
	"testing"
	"github.com/ethereum/go-ethereum/log"
	"reflect"
	"fmt"
)
type ControllerInterface interface {
	Init(action string, method string)
}

type Controller struct {
	Action string
	Method string
	Tag string `json:"tag"`
}

func (c *Controller) Init(action string, method string){
	c.Action = action
	c.Method = method
}

func TestController(t *testing.T){
	runController := &Controller{
		Action:"Run1",
		Method:"GET",
	}

	//Controller实现了ControllerInterface方法,因此它就实现了ControllerInterface接口
	var i ControllerInterface
	i = runController

	// 得到实际的值,通过v我们获取存储在里面的值,还可以去改变值
	v := reflect.ValueOf(i)
	fmt.Println("value:",v)

	vs := v.Elem()
	fmt.Println(vs.Field(0).String())
	fmt.Println(vs.Field(1).String())

	// 得到类型的元数据,通过t我们能获取类型定义里面的所有元素
	x := reflect.TypeOf(i)
	fmt.Println("type:",x)
	fmt.Println("type:",x.String())

}

func TestVar1(t *testing.T){
	type MyInt int
	var x MyInt = 7
	v := reflect.ValueOf(x)
	log.Info(string(v.Int()))
}
