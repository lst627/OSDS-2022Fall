package main

import (
	"os"
	"image/jpeg"
	"github.com/nfnt/resize"
	"log"
	//"sync"
	"bytes"
	"net"
	"net/rpc"
	//"net/http"
)

type Request struct {
	Key string
	Width int 
	Height int
}

type Reply []byte

type GetImg struct {
}

var Pool chan struct{}
var Count int

func (t *GetImg) GetSingleImg(request *Request, reply *Reply) error {
	// convert key to path !!!!!
	path := "/ImageNet/"+request.Key[0:2]+"/"+request.Key[2:4]+"/"+request.Key[4:6]+".JPEG"
	file, err := os.Open(path)
	if err != nil {
		log.Println("Load picture failed", request.Key)
		log.Fatal(err)
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Println("Decode picture failed", request.Key)
		log.Fatal(err)
	}
	file.Close()

	m := resize.Resize(uint(request.Width), uint(request.Height), img, resize.Lanczos3)
	buf := new(bytes.Buffer)
	err1 := jpeg.Encode(buf, m, nil)
	if err1 != nil {
		log.Println("Encode picture failed", request.Key)
		log.Fatal(err1)
	}
	send := buf.Bytes()

	*reply = send
	return nil
}

func (t *GetImg) GetQueueLength(request *Request, reply *int) error {
	*reply = Count
	return nil
}

func main() {
	num_workers := 10           // number of working threads
	Pool = make(chan struct{}, num_workers)
	getimg := new(GetImg)
	err := rpc.Register(getimg)
	if err != nil {
		log.Println("Service register failed")
		log.Fatal(err)
	} 
	//rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":6005")
	if err != nil {
		log.Println("Server error")
		log.Fatal(err)
	}
	log.Println("Ready to work!")
	for {
		conn, err := l.Accept()
		Count += 1
		if err != nil {
			log.Println("Accept error")
			log.Fatal(err)
		}
		log.Println("Request accepted")
		go func() {
			Pool <- struct{}{}
			Count -= 1
			defer func() { <- Pool}()
			rpc.ServeConn(conn) 
		}()
	}
	// go http.Serve(l, nil)

	log.Println("shutting down")
}