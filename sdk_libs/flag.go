package main

import (
	"flag"
	"fmt"
	"os"
)

type test string

//决定如何处理多个参数：-t a -t b -t c ，如果不实现Set方法，默认实现是取最后一个值
func (t *test) Set(s string) error {	fmt.Println("new value:" + s)
	*t = test(string(*t) + s)//多个参数拼接在一起
	return nil
}
func (t *test) String() string {
	return string(*t)
}
//自定义获取时的处理
func (t *test) Get() interface{} {
	tmp := "处理一下参数值：" + string(*t)
	return test(tmp)
}
var fs = flag.NewFlagSet("test", flag.ExitOnError)

func main() {
	var t test = test("default-value")
	fs.Var(&t, "t", "show info")
	fs.Parse(os.Args[1:])
	s := fs.Lookup("t").Value.String()
	fmt.Println("param:", s)
	tmp := fs.Lookup("t").Value.(flag.Getter).Get()
	fmt.Println("param:", tmp)
	a, b := fs.Lookup("t").Value.(flag.Getter).Get().(test)//Get返回值为interface{}类型
	fmt.Println("param:", a, b)

}