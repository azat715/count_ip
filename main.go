package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sync"

	parser "example.com/count_ip/internal/parser"
)

func main() {
	var wg sync.WaitGroup

	go func() {
		fmt.Println(http.ListenAndServe("127.0.0.1:6060", nil))
	}()

	wg.Add(1)
	go func() {
		res, err := parser.Parser(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)

		defer wg.Done()
	}()

	wg.Wait()
}
