package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {

	if len(os.Args) < 2 || len(os.Args) > 4 {
		usage(nil)
	}

	retry := 10
	delay := 1.0

	addr := os.Args[1]

	if len(os.Args) > 2 {
		i, err := strconv.Atoi(os.Args[2])
		if err != nil {
			usage(err)
		}
		retry = i
	}

	if len(os.Args) > 3 {
		f, err := strconv.ParseFloat(os.Args[3], 64)
		if err != nil {
			usage(err)
		}
		delay = f
	}

	for {
		_, err := net.Dial("tcp", addr)
		if err == nil {
			// success!
			break
		}
		retry--
		if retry == 0 {
			os.Exit(1)
		}
		log.Println("will retry in", delay, "sec (remaining:", strconv.Itoa(retry)+")")
		time.Sleep(time.Duration(delay*1000) * time.Millisecond)
	}
}

func usage(err error) {
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	fmt.Println("Usage: wait-for-service address:port [retries [delay]]")
	fmt.Println("Retries 10 times by default with 1 second delays")
	os.Exit(1)
}
