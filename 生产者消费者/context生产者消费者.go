package 生产者消费者

import (
	"context"
	"fmt"
	"time"
)

func Producer1(ctx context.Context, ch chan int) {
	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			return
		case ch <- i:
		}
	}
}

//消费者
func Consumer1(ctx context.Context, ch chan int) {
	for value := range ch {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(value)
		}
	}
}

func main() {
	ch := make(chan int, 60)
	ctx, cancel := context.WithCancel(context.Background()) //通过 context 来退出goroutine

	go Producer1(ctx, ch)
	go Consumer1(ctx, ch)

	time.Sleep(2 * time.Second)
	cancel()
}
