package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func handleConnection(c net.Conn) {
	FILENAME := ""
	var file *os.File
	reader := bufio.NewReader(c)
	start := time.Now()
	timePoint := time.Now()
	var totalWrite float64
	write := 0
	defer file.Close()
	defer c.Close()
	for {
		netData, _, err := reader.ReadLine()
		if err != nil {
			fmt.Printf("Final speed %.1f Mbyte/sec\n", totalWrite/time.Since(start).Seconds()/(1024*1024))
			return
		}
		if FILENAME == "" {
			FILENAME = strings.TrimSpace(string(netData)) + "copy"
			_, err := os.Create(FILENAME)
			if err != nil {
				log.Fatal(err)
			}
			file, _ = os.OpenFile(FILENAME, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			continue
		}
		writeNow, err := file.Write([]byte(append(netData, []byte("\n")...)))
		write += writeNow
		totalWrite += float64(writeNow)
		if err != nil {
			log.Fatal(err)
			return
		}
		duration := time.Since(timePoint).Seconds()
		if duration > 3 {
			fmt.Printf("Current speed %.1f Mbyte/sec\n", (float64(write)/duration)/(1024*1024))
			fmt.Printf("Total speed %.1f Mbyte/sec\n", totalWrite/time.Since(start).Seconds()/(1024*1024))
			timePoint = time.Now()
		}
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a address number!")
		return
	}

	ADDRESS := arguments[1]
	l, err := net.Listen("tcp4", ADDRESS)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		fmt.Printf("New client : %v\n", l.Addr())
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}
