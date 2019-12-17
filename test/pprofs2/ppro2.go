package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main()  {

	go func() {
		log.Println(http.ListenAndServe("localhost:8080",nil))
	}()
	go hello()
	select{
	}
}

func hello()  {
	for{
		go func() {
			fmt.Println("hello")
		}()
		time.Sleep(time.Millisecond*1)
	}
}
