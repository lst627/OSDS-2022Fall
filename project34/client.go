package main

import (
	"log"
	"sync"
	"time"
	"net/rpc"
	//"net"
	//"flag"
	// "fmt"
	//"crypto/sha256"
	//"math"
	//"bytes"
	//"strings"
)

type Block struct{
	Index 		int
	Timestamp 	time.Time
	Data 		map[string]interface{}
	PrevHash 	string
	Hash 		string
	Nonce		int
}

type Request struct {
	From string
	To string 
	Amount float64
	Index int
}

var miners map[string]*rpc.Client

func CallMiner(minerid string, request Request, wg *sync.WaitGroup) {
	var reply bool
	miner := miners[minerid]
	divCall := miner.Go("MinerHandler.GetTransaction", &request, &reply, nil)
	replyCall := <-divCall.Done
	if replyCall.Error != nil {
		log.Println("Error: ", replyCall.Error)
		delete(miners, minerid)
	}
	wg.Done()
}

func NewTranx(index int) {
	from := "Alice"
	to := "Bob"
	amount := float64(index)

	request := Request{from, to, amount, index}
	var wg = sync.WaitGroup{}
	wg.Add(len(miners))
	for id, _ := range miners {
		go CallMiner(id, request, &wg)
	}
	wg.Wait()
}

func main() {    
	num_tranx := 10
	serverAddresses := [5]string{"10.1.0.11", "10.1.0.12", "10.1.0.13", "10.1.0.14", "10.1.0.15"}   //"122.200.68.26"
	portNumbers := "6353" 

	miners = make(map[string]*rpc.Client)

	// Check miners
	// for _, miner := range miners {
	// 	var reply string
	// 	request := "Hi"
	// 	divCall := miner.Go("MinerHandler.GetReply", &request, &reply, nil)
	// 	replyCall := <-divCall.Done
	// 	if replyCall.Error != nil {
	// 		log.Fatal("Error: ", replyCall.Error)
	// 	}
	// 	log.Println(reply)
	// }

	start := time.Now()

	for i := 0; i < num_tranx; i++ {
		log.Println(i)

		for _, address := range serverAddresses {
			if _, ok := miners[address]; !ok {
				miner, err := rpc.Dial("tcp", address + ":" + portNumbers)
				if err != nil {
					//log.Println(address+ ":" + portNumbers)
					//log.Println("Miner dial failed")
					//log.Fatal(err)
				} else { 
					defer miner.Close() 
					miners[address] = miner
					log.Println("Connected to " + address+ ":" + portNumbers)
				}
			}
		}

		NewTranx(i)
		//time.Sleep(time.Duration(5) * time.Second)
	}

	cost := time.Since(start)
	log.Printf("cost=[%s]", cost)
}