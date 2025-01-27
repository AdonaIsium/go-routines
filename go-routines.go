package main

import (
	"fmt"
	"sync"
)

func reader(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for{
		val, ok := <-ch
		if !ok {
			fmt.Printf("%d 다닸어요\n", id)
			return
		}
		fmt.Printf("%d는 %d바닸어요\n", id, val)
	}

	
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(2)

	go reader(1, ch, &wg)
	go reader(2, ch, &wg)
	go reader(3, ch, &wg)
	go reader(4, ch, &wg)
	go reader(5, ch, &wg)
	go reader(6, ch, &wg)
	go reader(7, ch, &wg)
	go reader(8, ch, &wg)
	go reader(9, ch, &wg)
	go reader(10, ch, &wg)
	go reader(11, ch, &wg)
	go reader(12, ch, &wg)

	for i := 0; i < 1000; i++ {
		ch <- i
	}


	close(ch)

	wg.Wait()
}