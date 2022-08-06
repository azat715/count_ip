package main

import (
	"fmt"
	"io"
	"log"
	"os"

	bitmap "example.com/count_ip/internal/bit"
	iptoint "example.com/count_ip/internal/ip_to_int"
	readfile "example.com/count_ip/internal/read_file"
)

func main() {
	f, closer, err := readfile.New(os.Args[1])
	if err != nil {
		log.Panic("open file:", err)
	}
	defer closer()
	b := bitmap.New()
	for {
		line, readErr := f.Readline()
		if (readErr != nil) && (readErr != io.EOF) {
			log.Println("read file:", readErr)
		}
		i, err := iptoint.Convert(line)
		if err != nil {
			log.Println("convert ip:", err)
		}
		b.Set(i)
		if readErr == io.EOF {
			break
		}
	}
	fmt.Println("uniq ip:", b.Count())
}
