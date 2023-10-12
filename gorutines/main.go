package main

import (
	"fmt"
	"sync"
)

type Data struct {
	Title  string
	Amount float64
	Count  int
}

func CreateGoroutine() {
	wg := sync.WaitGroup{}
	c := make(chan Data, 1)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go DoSomethingMultithreading(&c, i, &wg)

	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for d := range c {
		fmt.Println(d)
	}
}

func DoSomethingMultithreading(c *chan Data, index int, wg *sync.WaitGroup) {
	defer wg.Done()

	d := Data{
		Title:  fmt.Sprintf("test %d", index),
		Amount: float64(index),
		Count:  index,
	}

	*c <- d
}

func main() {
	CreateGoroutine()
}
