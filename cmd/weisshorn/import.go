package main

import (
	"fmt"
	wh "github.com/gophergala2016/weisshorn"
	"strconv"
	"time"
)

func Import() {
	fmt.Printf("weisshorn client\n")

	sink := make(chan string, 1)
	//defer close(sink)
	producer := wh.NewLineProducer(sink)
end:
	for {
		select {
		case x, ok := <-sink:
			fmt.Println(">> " + x)
			fmt.Println("\t>> " + strconv.FormatBool(ok))
			if !ok {
				break end
			}
		case <-time.After(time.Second * 1):
			fmt.Println("timeout")
			break end
		}
	}

	producer.Stop()
	fmt.Println("bye bye")
}
