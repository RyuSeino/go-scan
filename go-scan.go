package main

import (
	"fmt"
	"flag"
	"net"
	"strconv"
	"time"
	"sync"
)

func scan(host string, i int, wg *sync.WaitGroup) {

	defer wg.Done()

	port := strconv.Itoa(i)
	conn, err := net.DialTimeout("tcp", host + ":" + port , 5 * time.Second)
	if err != nil {
		//handle error
	} else {
		fmt.Println(port + " is opened")
		conn.Close()
	}

}


func main() {

	h := flag.String("h", "localhost", "target host you scan")
	flag.Parse()
	fmt.Println(*h)

	wg := new(sync.WaitGroup)

	for i := 1; i < 1025; i++ {
		wg.Add(1)
		go scan(*h, i, wg)
	}

	wg.Wait()
}

