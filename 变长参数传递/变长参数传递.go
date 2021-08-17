package 变长参数传递

import "fmt"

func main() {
	x := min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	slice := []int{7, 9, 3, 5, 1}
	//如果参数被存储在一个 slice 类型的变量 slice 中，则可以通过 slice... 的形式来传递参数调用变参函数。
	x = min(slice...)
	fmt.Printf("The minimum in the slice is: %d", x)
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func F1(s ...string) {
	//变长参数可以作为对应类型的 slice 进行二次传递。
	F2(s...)
	F3(s)
}

func F2(s ...string) {}
func F3(s []string)  {}

func typeCheck(values ...interface{}) {
	for _, value := range values {
		switch v := value.(type) {
		case int:
			fmt.Println(v)
		case string:
			fmt.Println(v)
		case bool:
			fmt.Println(v)
		default:
			fmt.Println(v)
		}
	}

}
