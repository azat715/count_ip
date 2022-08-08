package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sync"

	bitmap "example.com/count_ip/internal/bit"
	iptoint "example.com/count_ip/internal/ip_to_int"
	readfile "example.com/count_ip/internal/read_file"
)

var bitArr = bitmap.New()

func count(s string) error {
	i, err := iptoint.Convert(s)
	if err != nil {
		return err
	}
	bitArr.Set(i)
	return nil
}

func main() {
	var wg sync.WaitGroup

	go func() {
		fmt.Println(http.ListenAndServe("127.0.0.1:6060", nil))
	}()

	wg.Add(1)
	go func() {
		f, closer, err := readfile.New(os.Args[1])
		if err != nil {
			log.Panic("open file:", err)
		}
		defer closer()

		for {
			line, readErr := f.Readline()
			if (readErr != nil) && (readErr != io.EOF) {
				log.Println("read file:", readErr)
			}
			err := count(line)
			if err != nil {
				log.Println("convert ip:", err)
			}
			if readErr == io.EOF {
				break
			}
		}
		fmt.Println("uniq ip:", bitArr.Count())
		defer wg.Done()
	}()

	wg.Wait()
}
