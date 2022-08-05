package main

import (
	"fmt"
	"io"
	"log"
	"os"

	iptoint "example.com/count_ip/internal/ip_to_int"
	readfile "example.com/count_ip/internal/read_file"
)

func main() {
	f, closer, err := readfile.New(os.Args[1])
	if err != nil {
		log.Panic("open file:", err)
	}
	defer closer()
loop:
	for {
		line, err := f.Readline()
		switch err {
		case io.EOF:
			fmt.Println("End file")
			break loop
		case nil:
			break
		default:
			log.Panic("read file:", err)
		}
		i, err := iptoint.Convert(line)
		if err != nil {
			log.Panic("convert ip:", err)
		}
		fmt.Println(i)

	}
}
