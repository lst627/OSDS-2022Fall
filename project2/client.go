package main

import (
	"log"
	"sync"
	"time"
	"net/rpc"
	"flag"
	"./producer_client"
	"./consumer_client"
	//"fmt"
)


var wg = sync.WaitGroup{}

var (
	st int
	ed int
)

func main() {
	flag.IntVar(&st, "start", 0, "")
	flag.IntVar(&ed, "end", 0, "")
	flag.Parse()
	//fmt.Println(st,ed)

    csvpath := "./resize.csv"

	num_workers := 10              // number of working threads
	capacity := 100                // task queue capacity
	task_queue := make(chan string, capacity)

	serverAddress := "10.1.0.15" //  //"122.200.68.26"
	portNumber := "6004"  //"8015"
	client, err := rpc.Dial("tcp", serverAddress + ":" + portNumber)
	if err != nil {
		log.Println("Client dial failed")
		log.Fatal(err)
	}
	defer client.Close()

	wg.Add(num_workers+1)
	start := time.Now()

	go producer_client.AddTasks(task_queue, csvpath, st, ed, &wg)
	for i := 0; i < num_workers; i++ {
		go consumer_client.GetImage(task_queue, client, &wg)
	}

	wg.Wait()

	cost := time.Since(start)
	log.Printf("cost=[%s]", cost)
}