package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

var timeout uint
var wg sync.WaitGroup
var connectionClosed = false

func init() {
	flag.UintVar(&timeout, "timeout", 10, "connection timeout")
	flag.Parse()
}

func readRoutine(ctx context.Context, conn net.Conn, closeConnChan chan<- int) {
	defer wg.Done()
	scanner := bufio.NewScanner(conn)
OUTER:
	for {
		select {
		case <-ctx.Done():
			break OUTER
		default:
			if !scanner.Scan() {
				log.Printf("CONNECTION WITH SERVER CLOSED\n")
				if !connectionClosed {
					closeConnChan <- 0
				}
				break OUTER
			}
			text := scanner.Text()
			log.Printf("From server: %s", text)
		}
	}
	log.Printf("reading from server finished")
}

func writeRoutine(ctx context.Context, conn net.Conn, closeConnChan chan<- int) {
	defer wg.Done()
	scanner := bufio.NewScanner(os.Stdin)
OUTER:
	for {
		select {
		case <-ctx.Done():
			break OUTER
		default:
			if !scanner.Scan() {
				log.Printf("GOT EOF")
				log.Printf("closing connection with server...")
				closeConnChan <- 1
				break OUTER
			}
			str := scanner.Text()

			log.Printf("To server %v\n", str)

			conn.Write([]byte(fmt.Sprintf("%s\n", str)))
		}

	}
	log.Printf("writing to server finished")
}

func main() {
	var host string
	var port string

	if len(os.Args) < 3 {
		log.Fatal("too few arguments")
	}

	wg.Add(2)

	if len(os.Args) == 4 {
		host = os.Args[2]
		port = os.Args[3]
	} else {
		host = os.Args[1]
		port = os.Args[2]
	}

	addr := host + ":" + port

	dialer := &net.Dialer{}
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	conn, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}

	closeConnChan := make(chan int, 1)

	go readRoutine(ctx, conn, closeConnChan)
	go writeRoutine(ctx, conn, closeConnChan)

	select {
	case res := <-closeConnChan:
		if res == 0 {
			log.Printf("please press enter to finish client")
		}

		connectionClosed = true
		cancel()
		conn.Close()
		close(closeConnChan)
	}

	wg.Wait()
}
