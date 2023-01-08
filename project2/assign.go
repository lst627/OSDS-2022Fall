package main

import (
	"os/exec"
	"log"
	"fmt"
	"os"
	"strconv"
)

func main() {
	num_processes := 10        // Number of forked processes
	size := 2000 / num_processes
	var start int 
	var end int
	for i := 0; i < num_processes; i++ {
		start = i*size 
		end = (i+1)*size
		if i == (num_processes - 1) {end = 2000}
		cmd := exec.Command("./client", "-start", fmt.Sprintf("%s", strconv.Itoa(start)),  "-end", fmt.Sprintf("%s", strconv.Itoa(end)))
		// For debugging mode:
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Start()
		if err != nil {
			log.Println("Fork error")
			log.Fatal(err)
		} //else { fmt.Println("Success!") }
	}
}