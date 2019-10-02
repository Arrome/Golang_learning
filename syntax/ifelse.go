package syntax

import "fmt"

func main() {

	// num 为if lock中局部变量
	if num := 9;num < 0 {
		fmt.Println(num,"is negative")
	} else if num < 10 {
		fmt.Println(num,"has 1 digit")
	} else {
		fmt.Println(num,"has multiple digits")
	}
	
}
