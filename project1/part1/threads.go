package main

import (
	"fmt"
	"sync"
	"time"
	"./concurrent_queue"
)

var wg = sync.WaitGroup{}

func Request (Q *concurrent_queue.ConcurrentQueue)() {
	var res int
	var str string = "a"
	res = Q.Enqueue(str)
	//fmt.Println(res)
	res = Q.Size()
	//fmt.Println(res)
	str, res = Q.Dequeue()
	//fmt.Println(res)
	res = Q.Capacity()
	//fmt.Println(res)
	res = Q.Size()
	fmt.Println(str, res)
	wg.Done()
}

func main() {
	var load int = 10
	var capacity int = 7

    var Q *concurrent_queue.ConcurrentQueue = concurrent_queue.NewQueue();
	Q.Init(capacity)
	wg.Add(load)
    
	start := time.Now()
	for i := 0; i < load; i++ {
		go Request(Q)
	}
	wg.Wait()
	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}