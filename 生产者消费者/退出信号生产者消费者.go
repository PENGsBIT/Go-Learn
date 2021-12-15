package 生产者消费者

import (
	"fmt"
	"time"
)

func Producer(ch chan int, quit chan bool) {
	for i := 1; ; i++ {
		select {
		case <-quit:
			return
		case ch <- i:
		}
	}
}

//消费者
func Consumer(ch chan int, quit chan bool) {
	for value := range ch {
		select {
		case <-quit:
			return
		default:
			fmt.Println(value)
		}
	}
}

func main() {
	ch := make(chan int, 60)
	quit := make(chan bool) //通过增加 一个退出标签 来确定退出

	go Producer(ch, quit)
	go Consumer(ch, quit)

	time.Sleep(2 * time.Second)
	quit <- true
}
