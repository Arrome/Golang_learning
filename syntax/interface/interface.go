package main
/**
	类型，抽象类型<br>
	接口的实现是隐式的，不需要声明，只需要实现方法

	* 接口嵌套
	* 空接口 ：可以传任意类型
		使用场景：map的value值类型
 */
import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width,height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64  {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return  math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

//接口类型，定义实现了需要使用方法的
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width:3,height:4}
	c := circle{radius:5}
	measure(r)
	measure(c)
}
