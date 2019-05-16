package main

import "fmt"

type rec struct {
	width,height int
}

func (r *rec) area() int {
	return r.width * r.height
}

func main() {
	r := rec{width:10,height:5}
	fmt.Println("area: ",r.area())
}
