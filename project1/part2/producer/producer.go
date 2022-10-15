package producer

import (
	"log"
	"io/ioutil"
	"sync"
	// "fmt"
)

func AddTasks (Q chan string, path_from string, wg *sync.WaitGroup)() {
	fileInfoList, err := ioutil.ReadDir(path_from)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(len(fileInfoList)) 
	for i := range fileInfoList {
		// fmt.Println(fileInfoList[i].Name())
		Q<-fileInfoList[i].Name()
	}
	Q<-"Finish"
	wg.Done()
}