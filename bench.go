package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

var host string
var nclient int

func init() {
	flag.StringVar(&host, "host", "localhost:5000", "host name")
	flag.IntVar(&nclient, "n", 10, "number of clients")
}

func client() {
	addr, err := net.ResolveTCPAddr("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer conn.Close()
		buf := make([]byte, 256)
		for {
			n, err := conn.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err)
				break
			}
			log.Println(string(buf[:n]))
		}
	}()

	for {
		_, err := conn.Write([]byte("Hello"))
		if err != nil {
			conn.CloseWrite()
			break
		}
		time.Sleep(time.Millisecond * 10)
	}
}

func main() {
	flag.Parse()
	for i := 0; i < nclient; i++ {
		go client()
	}

	for {
		time.Sleep(time.Second)
	}
}
