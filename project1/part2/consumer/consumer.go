package consumer

import (
	"os"
	"log"
	"image/jpeg"
	"github.com/nfnt/resize"
	"fmt"
	"sync"
)

func Work (Q chan string, path_from string, path_to string, wg *sync.WaitGroup)() {
	var ok bool
	var filename string
	filename, ok = <-Q //"resize_demo.JPEG"
	for {
		if !ok { break }
		if filename == "Finish" { 
			close(Q)
			break 
		}
		path_name := fmt.Sprintf("%s/%s", path_from, filename)
		// fmt.Println(path_name)
		file, err := os.Open(path_name)
		if err != nil {
			fmt.Println("Load picture failed", filename)
			log.Fatal(err)
		}
		img, err := jpeg.Decode(file)
		if err != nil {
			fmt.Println("Decode picture failed", filename)
			log.Fatal(err)
		}
		file.Close()

		m := resize.Resize(128, 128, img, resize.Lanczos3)
		
		out, err := os.Create(fmt.Sprintf("%s/%s", path_to, filename))       //("resize_demo_resized.JPEG")
		if err != nil {
			fmt.Println("Write picture failed", filename)
			log.Fatal(err)
		}
		defer out.Close()

		jpeg.Encode(out, m, nil)

		filename, ok = <-Q //"resize_demo.JPEG"
	} //for
	wg.Done()
}