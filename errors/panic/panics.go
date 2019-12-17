package main

import "fmt"

/**
panic 表示非常严重的不可恢复的错误
*/

func panices()  {

	defer func() {
		if n := recover(); n != nil{
			fmt.Printf("%s\n",n)
		}
	}()

	panic("抛出一个不可恢复且特别严重的错误")
}

func main()  {

	panices()
	fmt.Println("hello world")
}
