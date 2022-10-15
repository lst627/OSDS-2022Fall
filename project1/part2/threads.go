package main

import (
	"os"
	"fmt"
	"log"
	"sync"
	"time"
	"./consumer"
	"./producer"
)

var wg = sync.WaitGroup{}

func MakeFolder(path_to string) {
	_, err := os.Stat(path_to)
	if os.IsNotExist(err) {
		err = os.Mkdir(path_to, os.ModePerm)
		if err != nil {
			fmt.Printf("Create folder failed -> %v\n", err)
			log.Fatal(err)
		}
	} else if err == nil{
		err = os.RemoveAll(path_to)
		if err != nil {
			fmt.Printf("Remove existing folder failed -> %v\n", err)
			log.Fatal(err)
		}
		err = os.Mkdir(path_to, os.ModePerm)
		if err != nil {
			fmt.Printf("Create folder failed -> %v\n", err)
			log.Fatal(err)
		}
	} else {
		fmt.Printf("Unknown error -> %v\n", err)
		log.Fatal(err)
	}
}

func main() {
	var num_workers int = 50
	// var load int = 10000
	var capacity int = 100
	path_from :=  "./tiny-imagenet-200/test/images" // "../test_image"
	path_to := "./tiny-imagenet-200/test/images_resize" //"../test_image_resize" 
	MakeFolder(path_to)
	task_queue := make(chan string, capacity)
	wg.Add(num_workers + 1)
    
	start := time.Now()
	go producer.AddTasks(task_queue, path_from, &wg)
	for i := 0; i < num_workers; i++ {
		go consumer.Work(task_queue, path_from, path_to, &wg)
	}
	wg.Wait()
	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}