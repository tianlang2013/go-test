package study

import (
"testing"
"fmt"
"math"
)

func TestMap(t *testing.T)  {

	fmt.Println("study")
}


func TestArray(t *testing.T){
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}

func TestSlice(t *testing.T){
	// 不想数组，slice 的类型仅有它所包含的元素决定（不像
	// 数组中还需要元素的个数）。要创建一个长度非零的空
	// slice，需要使用内建的方法 `make`。这里我们创建了一
	// 个长度为3的 `string` 类型 slice（初始化为零值）。
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// 我们可以和数组一起设置和得到值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// 如你所料，`len` 返回 slice 的长度
	fmt.Println("len:", len(s))

	// 作为基本操作的补充，slice 支持比数组更多的操作。
	// 其中一个是内建的 `append`，它返回一个包含了一个
	// 或者多个新值的 slice。注意我们接受返回由 append
	// 返回的新的 slice 值。
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slice 也可以被 `copy`。这里我们创建一个空的和 `s` 有
	// 相同长度的 slice `c`，并且将 `s` 复制给 `c`。
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slice 支持通过 `slice[low:high]` 语法进行“切片”操
	// 作。例如，这里得到一个包含元素 `s[2]`, `s[3]`,
	// `s[4]` 的 slice。
	l := s[2:5]
	fmt.Println("sl1:", l)

	// 这个 slice 从 `s[0]` 到（但是不包含）`s[5]`。
	l = s[:5]
	fmt.Println("sl2:", l)

	// 这个 slice 从（包含）`s[2]` 到 slice 的后一个值。
	l = s[2:]
	fmt.Println("sl3:", l)

	// 我们可以在一行代码中申明并初始化一个 slice 变量。

	x := []byte{'a','b'}
	fmt.Println("dcl:", x)
	fmt.Println("dcl string"  , string(x))
	// Slice 可以组成多维数据结构。内部的 slice 长度可以不
	// 同，这和多位数组不同。
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	ss := "abcd"
	fmt.Println("string slice" , ss[1:3])

}
type geometry interface {
	area() float64
	perim() float64
}
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}
func (r *rect) area() float64 {
	return r.width * r.height
}
func (r *rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func(r *rect) x(){
	return
}
func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c *circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func TestInterface( t *testing.T)  {
	r := &rect{width: 3, height: 4}
	c := &circle{radius: 5}

	measure(r)
	measure(c)
}

func TestChain(t *testing.T)  {
	c := make(chan int ,10 )

	go func(){
		fmt.Println("start")
		for i:=0 ;i<=10 ; i++  {
			c <- i
		}
	}()

	go func() {
		for{
			select {
			case cc :=<- c:
				fmt.Println(cc)
			}

		}
	}()
}


