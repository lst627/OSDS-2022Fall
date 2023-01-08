package producer_client

import (
	"log"
	"sync"
	"encoding/csv"
	"os"
	"io"
)

func AddTasks (Q chan string, csvpath string, st int, ed int, wg *sync.WaitGroup)() {
	f, err := os.Open(csvpath)
	if err != nil {
		log.Println("Open csv failed")
		log.Fatal(err)
	}
	defer f.Close()
	r := csv.NewReader(f)

	for i := 0; i < 2000; i++ {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Println("Read csv failed")
			log.Fatal(err)
		}
		if err == io.EOF {
			break
		}
		//log.Println(row)
		if i>=st { Q<-row[0] }
		if i>=ed { break }
	}
	Q<-"Finish"

	wg.Done()
}