package 值传递与引用传递

import "fmt"

func main() {
	var args int64 = 1
	modifiedNumber(args) // args就是实际参数
	fmt.Printf("实际参数的地址 %p\n", &args)
	fmt.Printf("改动后的值是  %d\n", args)
	//形参地址 0xc0000b4010
	//实际参数的地址 0xc0000b4008
	//改动后的值是  1
	addr := &args
	fmt.Printf("原始指针的内存地址是 %p\n", addr)
	fmt.Printf("指针变量addr存放的地址 %p\n", &addr)
	modifiedNumberPtr(addr)
	fmt.Printf("改动后的值是  %d\n", args)
	//原始指针的内存地址是 0xc0000b4008
	//指针变量addr存放的地址 0xc0000ae018
	//形参地址 0xc0000ae028
	//改动后的值是  10

}

func modifiedNumber(args int64) { //这里定义的args就是形式参数
	fmt.Printf("形参地址 %p \n", &args)
	args = 10
}

func modifiedNumberPtr(addr *int64) { //这里定义的args就是形式参数
	fmt.Printf("形参地址 %p \n", &addr)
	*addr = 10
}

//以确定go就是值传递，因为我们在modifieNumber方法中打印出来的内存地址发生了改变，
//但是说好的go中只有值传递呢，
//为什么chan、map、slice类型传递却可以改变其中的值呢？
//我们看一下slice底层结构：
// //runtime/slice.go
// type slice struct {
//    array unsafe.Pointer
//    len   int
//    cap   int
// }
//他的第一个元素是一个指针类型，这个指针指向的是底层数组的第一个元素。
//为什么slice也是值传递。之所以对于引用类型的传递可以修改原内容的数据，这是因为在底层默认使用该引用类型的指针进行传递，
//但也是使用指针的副本，依旧是值传递。所以slice传递的就是第一个元素的指针的副本，因为fmt.printf缘故造成了打印的地址一样，给人一种混淆的感觉。
//map
//源码入手，看一看什么原理：
//func makemap(t *maptype, hint int, h *hmap) *hmap {
//从以上源码，我们可以看出，使用make函数返回的是一个hmap类型的指针*hmap。
//举个例子，我们的func modifiedAge(person map[string]int)函数，其实就等于func modifiedAge(person *hmap）

//实际上在作为传递参数时还是使用了指针的副本进行传递，属于值传递。
//go就是值传递，可以确认的是Go语言中所有的传参都是值传递（传值），都是一个副本，一个拷贝。
//因为拷贝的内容有时候是非引用类型（int、string、struct等这些），这样就在函数中就无法修改原内容数据；
//有的是引用类型（指针、map、slice、chan等这些），这样就可以修改原内容数据。
