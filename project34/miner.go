package main

import (
	"log"
	"sync"
	"time"
	"net/rpc"
	"flag"
	"fmt"
	"crypto/sha256"
	//"math"
	//"bytes"
	"strings"
	//"os"
	"strconv"
	"net"
	"encoding/json"
)

type Block struct{
	Index 		int
	Timestamp 	time.Time
	Data 		map[string]interface{}
	PrevHash 	string
	Hash 		string
	Nonce		int
	Miner		string
}

type Blockchain struct {
	GenesisBlock Block
	Chain 		[]Block
	Difficulty 	int
}

type Request struct {
	From string
	To string 
	Amount float64
	Index int
}

type NewPeer struct {
	ServerAddress string
	ServerPortnumber string
}

var peers map[string]*rpc.Client
var selfAddress string

var B Blockchain
var Q chan Block
var num_tranx int

func (b *Blockchain) CreateBlockchain(difficulty int) {
	genesisBlock := Block{
			Hash:      "0",
			Timestamp: time.Now(),
			Index:    0,
	}
	b.GenesisBlock = genesisBlock
	b.Chain = []Block{genesisBlock}
	b.Difficulty = difficulty
}

func (b *Block) CalculateHash() string {
	data, _ := json.Marshal(b.Data)
	blockData := b.PrevHash + string(data) + b.Timestamp.String() + strconv.Itoa(b.Nonce)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) Mine(difficulty int) bool {
	for !strings.HasPrefix(b.Hash, strings.Repeat("0", difficulty)) {
			if num_tranx >= b.Index {
				return false
			}
			b.Nonce ++
			b.Hash = b.CalculateHash()
	}
	log.Println("Mining block", b.Index, "succeeded!")
	return true
}

func (b *Blockchain) MineBlock(from, to string, amount float64, index int)(reply Block) {
	data := map[string]interface{}{
			"from":   from,
			"to":     to,
			"amount": amount,
	}
	lastBlock := b.Chain[len(b.Chain)-1]
	newBlock := Block{
			Data:         data,
			PrevHash: lastBlock.Hash,
			Timestamp:    time.Now().Round(0),
			Index:		index,
			Miner:      selfAddress,
	}
	if newBlock.Mine(b.Difficulty) {
		return newBlock
	}
	newBlock.Index = -1
	return newBlock
}

func (b *Blockchain) IsValid() bool {
	length := len(b.Chain)
	for i := 1 ; i <= length - 2; i++  {
			previousBlock := b.Chain[i]
			currentBlock := b.Chain[i+1]
			if currentBlock.Hash != currentBlock.CalculateHash() || currentBlock.PrevHash != previousBlock.Hash {
					log.Println("Error. Invalid chain.")
					return false
			}
	}
	return true
}

func (b *Blockchain) AddBlock(newBlock Block) bool {
	if !b.IsValid() {
		b.Chain = []Block{b.GenesisBlock}
		return false
	}
	b.Chain = append(b.Chain, newBlock)
	return true
}

func (b *Blockchain) Tail() int {
	return b.Chain[len(b.Chain)-1].Index
}


func ReceiveBlock() {
	var r Block
	var ok bool
	r, ok = <-Q 
	for {
		if !ok { break }
		// if r.Data["from"] == "Finish" { 
		// 	close(Q)
		// 	break 
		// }
		if r.Index > num_tranx {
			num_tranx = r.Index
			log.Println("Adding block", num_tranx, "!")
			B.AddBlock(r)
		}
		r, ok = <-Q 
	}
}

func Login(initserverAddresses [5]string, clientAddress string, portnumber string) {
	for _, address := range initserverAddresses {
		// if address == selfAddress { continue }
		miner, err := rpc.Dial("tcp", address + ":" + portnumber)
		if err != nil {
			log.Println(address+ ":" + portnumber + " Miner dial failed")
			//log.Fatal(err)
		} else { 
			peers[address] = miner
		}
	}
	var reply bool
	var send NewPeer
	send.ServerAddress = selfAddress
	send.ServerPortnumber = portnumber
	for _, miner := range peers {
		miner.Go("MinerHandler.GetNewPeer", &send, &reply, nil)
	}
}

func BroadcastBlock(minerid string, send Block, wg *sync.WaitGroup) {
	var reply bool
	miner := peers[minerid]
	divCall := miner.Go("MinerHandler.GetBlock", &send, &reply, nil)
	replyCall := <-divCall.Done
	if replyCall.Error != nil {
		log.Println("Error: ", replyCall.Error)
		delete(peers, minerid)
	}
	wg.Done()
}

type MinerHandler struct {
}

func (g *MinerHandler) GetTransaction(request *Request, reply *bool) error {
	// log.Println("Received a block", request.Index)
	send := B.MineBlock(request.From, request.To, request.Amount, request.Index)
	*reply = true
	// Broadcast Block
	if send.Index >= 0 {
		var wg = sync.WaitGroup{}
		wg.Add(len(peers))
		for id, _ := range peers {
			go BroadcastBlock(id, send, &wg) 
		}
		wg.Wait()
	}
	return nil
}

func (g *MinerHandler) GetBlock(request *Block, reply *bool) error {
	Q <- *request
	*reply = true
	return nil
}

func (g *MinerHandler) GetReply(request *string, reply *string) error {
	*reply = selfAddress
	return nil
}

func (g *MinerHandler) GetNewPeer(request *NewPeer, reply *bool) error {
	miner, err := rpc.Dial("tcp", request.ServerAddress + ":" + request.ServerPortnumber)
	if err != nil {
		log.Println("Adding new peer failed!" + request.ServerAddress+ ":" + request.ServerPortnumber)
		//log.Fatal(err)
	} else { 
		peers[request.ServerAddress] = miner
		log.Println("Adding new peer succeed!" + request.ServerAddress+ ":" + request.ServerPortnumber)
	}
	return nil
}

func main() {
	var difficulty int
	var portnumber string
	Q = make(chan Block, 100)
	num_tranx = -1
	peers = make(map[string]*rpc.Client)
	
	flag.IntVar(&difficulty, "diff", 0, "")
	flag.StringVar(&portnumber, "port", "6353", "")
	flag.StringVar(&selfAddress, "addr", "10.1.0.11", "")
	flag.Parse()
	initserverAddresses := [5]string{"10.1.0.11", "10.1.0.12", "10.1.0.13", "10.1.0.14", "10.1.0.15"}   //"122.200.68.26"
	clientAddress := "10.1.0.16"
	
	B.CreateBlockchain(difficulty)
	gettranx := new(MinerHandler)
	err := rpc.Register(gettranx)
	if err != nil {
		log.Println("Service register failed")
		log.Fatal(err)
	} 
	//rpc.HandleHTTP()
	// Listen
	l, err := net.Listen("tcp", ":"+portnumber)
	if err != nil {
		log.Println("Miner error")
		log.Fatal(err)
	}
	// Login and find initial peers
	Login(initserverAddresses, clientAddress, portnumber)

	// Start to receive blocks
	go ReceiveBlock()
	log.Println("Ready to mine!")
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Connection error")
			log.Fatal(err)
		}
		log.Println("New connection established")
		go rpc.ServeConn(conn) 
	}
	// go http.Serve(l, nil)

	log.Println("Shutting down")
}