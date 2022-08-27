package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"text/tabwriter"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type cityServer map[string]string
type cityTime map[string]string

func main() {

	cityServers := make(cityServer)

	timeZones := os.Args[1:]
	cityTimes := make(cityTime)

	for _, v := range timeZones {
		sp := strings.Split(v, "=")
		cityServers[sp[0]] = sp[1]
	}

	var err error
	for i, v := range cityServers {
		cityTimes[i], err = requestLocal(v)
		check(err)
	}
	printTimezones(cityTimes)
}

func printTimezones(cities cityTime) {
	const format = "%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "City", "Time")
	fmt.Fprintf(tw, format, "-----", "-----")
	for i, v := range cities {
		fmt.Fprintf(tw, format, i, v)
	}
	tw.Flush()
}

func requestLocal(addr string) (string, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return "", err
	}
	clock, err := bufio.NewReader(conn).ReadString('\n')
	check(err)
	return clock, nil
}
