package 闭包

import (
	"fmt"
	"sync"
)

//在main函数中由于在匿名函数引用了外部变量i,因此匿名函数以闭包的形式访问i,
//当使用闭包方式访问某个局部变量时，该局部变量会常驻内存，访问速度会特别快，并且我们始终操作的都是这个变量的本身，而不是其副本，修改时也特别方便。
//但同时for循环中使用了goroutine,由于goroutine之间是并发执行的，再参考图2闭包调用的流程图，
//就会出现多个goroutine访问同一个内存中的变量，会出现“脏读”现象

var wg sync.WaitGroup

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		// 以匿名函数的形式开启goroutine
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

//为了解决这个问题，也很简单，我们只需让for循环时每个匿名函数绑定的不是外部变量i,而是i的副本，如何解决呢？
//之间用匿名函数传参的形式，由于go语言的函数传参都是值传递，这样就可以通过值传递来为每个goroutine的匿名函数复制出一个当前i的副本，
//所有的goroutine在同时执行时互不影响。
var wg1 sync.WaitGroup

func main1() {
	wg1.Add(10)
	for i := 0; i < 10; i++ {
		// 以匿名函数的形式开启goroutine
		go func(num int) {
			fmt.Println(num)
			wg1.Done()
		}(i)
	}
	wg1.Wait()
}
