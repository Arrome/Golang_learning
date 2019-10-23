package main

import "fmt"

type intValue int

const (
	Running intValue = iota
	Stopped
	Rebooting
	Terminated
)

//枚举类型，给常量增加一个String()方法，增加可读性
func (s intValue) String() string {
	switch s {
	case Running:
		return "Running"
	case Stopped:
		return "Stopped"
	case Rebooting:
		return "Rebooting"
	case Terminated:
		return "Terminated"
	default:
		return "Unknown"
	}
}

func main() {
	state := Running
	fmt.Println(state.String())
}
