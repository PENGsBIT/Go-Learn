package 闭包

import "fmt"

func incr() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

//调用这个函数会返回一个函数变量。
//
//i := incr()：通过把这个函数变量赋值给 i，i 就成为了一个闭包。
//
//所以 i 保存着对 x 的引用，可以想象 i 中有着一个指针指向 x 或 i 中有 x 的地址。
//
//由于 i 有着指向 x 的指针，所以可以修改 x，且保持着状态：
func main() {
	i := incr()
	fmt.Println(i()) // 1
	fmt.Println(i()) // 2
	fmt.Println(i()) // 3
	//x 逃逸了，它的生命周期没有随着它的作用域结束而结束。
	println(incr()()) // 1
	println(incr()()) // 1
	println(incr()()) // 1
	//因为这里调用了三次 incr()，返回了三个闭包，这三个闭包引用着三个不同的 x，它们的状态是各自独立的。
	//相当于 i1 := incr() i2 := incr() i3 := incr()
	//调用 i1(),i2(),i3()
}
