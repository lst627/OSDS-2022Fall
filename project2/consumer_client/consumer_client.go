package consumer_client

import (
	"os"
	"log"
	"image/jpeg"
	"sync"
	"net/rpc"
	//"strconv"
	"bytes"
	"image"
)

type Request struct {
	Key string
	Width int 
	Height int
}

type Reply []byte

func SaveImage(imgbyte []byte, key string)() {
	img, _, err := image.Decode(bytes.NewReader(imgbyte))
    if err != nil {
		log.Println("Decode image failed")
        log.Fatal(err)
    }

    out, _ := os.Create("./"+key+".jpg")
    defer out.Close()

    var opts jpeg.Options
    opts.Quality = 1

    err = jpeg.Encode(out, img, &opts)
    //jpeg.Encode(out, img, nil)
    if err != nil {
		log.Println("Save image failed")
        log.Fatal(err)
    }
}

func GetImage(Q chan string, client *rpc.Client, wg *sync.WaitGroup)() {
	var ok bool
	var key string
	var height, width int

	key, ok = <-Q 
	for {
		if !ok { break }
		if key == "Finish" { 
			close(Q)
			break 
		}
		//width, _ = strconv.Atoi(row[1])
		//height, _ = strconv.Atoi(row[2])
		width = 128
		height = 128
		request := Request{key, width, height}
		var reply Reply
		divCall := client.Go("GetImg.GetSingleImg", &request, &reply, nil)
		
		replyCall := <-divCall.Done
		
		if replyCall.Error != nil {
			log.Fatal("Error: ", replyCall.Error)
		}
		SaveImage(reply, key)
		key, ok = <-Q 
	}
	//log.Println("Succeed!", key)
	wg.Done()
}
