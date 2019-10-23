package main

import (
	"flag"
	"fmt"
	"os"
)

var cli = flag.CommandLine

func main() {
	fmt.Println(string("hello world!你好，世界"))
	var intValue int
	cli.IntVar(&intValue,"t",3,"int value")
	cli.Parse(os.Args)
	t := cli.Lookup("t").Value.(flag.Getter).Get()
	fmt.Println(intValue,t)
}
