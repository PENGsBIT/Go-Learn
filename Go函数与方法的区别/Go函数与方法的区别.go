package Go函数与方法的区别

import "fmt"

//我们定义了add就是一个函数，它的函数签名是func add(a, b int) int,没有接收者，直接定义在go的一个包之下，可以直接调用，
//比如例子中的main函数调用了add函数。例子中的这个函数名称是小写开头的add，所以它的作用域只属于所声明的包内使用，
//不能被其他包使用，如果我们把函数名以大写字母开头，该函数的作用域就大了，可以被其他包调用。这也是Go语言中大小写的用处，
//比如Java中就有专门的关键字来声明作用域private、protect、public等。如下例子中定义的Add方法就可以被其他包调用。
func main() {
	sum := add(1, 2)
	fmt.Println(sum)
}

func add(a, b int) int {
	return a + b
}

// 一个加法实现：返回a+b的值
func Add(a, b int) int {
	return a + b
}

// 方法的声明和函数类似，他们的区别是：方法在定义的时候，会在func和方法名之间增加一个参数，这个参数就是接收者，
//这样我们定义的这个方法就和接收者绑定在了一起，称之为这个接收者的方法。
type person struct {
	name string
}

//接受者p 类型person 方法名String 返回类型string
func (p person) String() string {
	return "the person name is " + p.name
}

//现在我们说，类型person有了一个String方法，现在我们看下如何使用它。
func main1() {
	p := person{name: "张三"}
	fmt.Println(p.String())
}

//Go语言里有两种类型的接收者：值接收者和指针接收者。我们上面的例子中，就是使用值类型接收者的示例。
//使用值类型接收者定义的方法，在调用的时候，使用的其实是值接收者的一个副本，所以对该值的任何操作，不会影响原来的类型变量。
func main2() {
	p := person{name: "张三"}
	p.valModify()           //值接收者，修改无效
	fmt.Println(p.String()) //打印出来的值还是张三，对其进行的修改无效
}

func (p person) valModify() {
	p.name = "李四"
}

//如果我们使用一个指针作为接收者，那么就会其作用了，因为指针接收者传递的是一个指向原值指针的副本，
//指针的副本，指向的还是原来类型的值，所以修改时，同时也会影响原来类型变量的值。
//我们可以简单的理解为值接收者使用的是值的副本来调用方法，而指针接收者使用实际的值来调用方法
func main4() {
	p := person{name: "张三"}
	p.ptrModify() //指针接收者，修改有效
	//(&p).modify() 同理于指针接收者，修改有效
	fmt.Println(p.String())
}

func (p *person) ptrModify() {
	p.name = "李四"
}
