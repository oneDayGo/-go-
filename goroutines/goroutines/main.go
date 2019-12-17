package main

import (
	"fmt"
	"time"
)

func main() {


	for i:=1;i<10;i++{
		//创建协程
		go func(i int) {
			fmt.Printf("%d",i)
		}(i)
	}

	time.Sleep(1e9)

}
