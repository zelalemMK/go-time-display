package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {

	//var nFlag = flag.Int("n", 1234, "help message for flag n")

	portString := flag.String("port", "8000", "Input port number")

	flag.Parse()

	if parsed := flag.Parsed(); !parsed {
		log.Fatal("Flag not parsed correctly.")
	}

	listener, err := net.Listen("tcp", "localhost:"+*portString)
	check(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	timezone, err := time.LoadLocation(os.Getenv("TZ"))
	check(err)

	defer c.Close()
	for {

		_, err := io.WriteString(c, time.Now().In(timezone).Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
