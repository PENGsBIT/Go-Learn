package 闭包

import (
	"fmt"
)

func main() {
	f := func(args string) {
		fmt.Println(args)
	}
	f("hello world") //hello world
	//或
	(func(args string) {
		fmt.Println(args)
	})("hello world") //hello world
	//或
	func(args string) {
		fmt.Println(args)
	}("hello world") //hello world
}
