package main

import "fmt"

type persion struct{
	name string
	age int
}

func main() {
	fmt.Println(persion{"Bob",20}) //{Bob 20}
	fmt.Println(persion{name: "Fred"}) //{Fred 0}
	fmt.Println(&persion{name:"Ann",age:40}) //&{Ann 40}

	s := persion{name: "Sean",age: 50}
	fmt.Println(s.name) // Sean

	sp := &s
	fmt.Println(sp.age)  // 50

	sp.age = 51
	fmt.Println(sp.age) //51
}
