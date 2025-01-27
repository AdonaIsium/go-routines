package main

import (
	"fmt"
	"sync"
)

func reader(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello ")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go reader(&wg)
	go reader(&wg)

	wg.Wait()
}