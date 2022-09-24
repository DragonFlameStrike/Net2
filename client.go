package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Println("Please provide host:port filename.")
		return
	}

	CONNECT := arguments[1]
	FILEPATH := arguments[2]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, _ := os.Open(FILEPATH)
	fmt.Fprintf(c, file.Name()+"\n")
	defer file.Close()
	reader := bufio.NewReader(file)
	written, err := io.Copy(c, reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(written)

}
