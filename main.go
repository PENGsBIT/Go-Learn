package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(time.Duration(rand.Intn(4)))
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done() // Done()用来表示goroutine已经完成了，减少一次计数器
}

func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait() // Wait()用来等待所有需要等待的goroutine完成。
	fmt.Println("All go routines finished executing")
}
