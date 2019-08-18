package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	for _, addressWithPort := range os.Args[1:] {
		wg.Add(1)
		go func(address string, waitgroup *sync.WaitGroup) {
			fmt.Println(address)
			listener, err := net.Listen("tcp", address)
			if err != nil {
				log.Fatal(err)
			}
			for {
				conn, err := listener.Accept()
				if err != nil {
					log.Print(err) // e.g., connection aborted
					continue
				}
				go handleConn(conn) // handle connections concurrently
			}
			waitgroup.Done()
		}(addressWithPort, &wg)
	}
	wg.Wait()
}
