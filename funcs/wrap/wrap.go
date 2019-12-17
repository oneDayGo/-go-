package main

import (
	"fmt"
	"time"
)

//闭包
func warps() func() int {
	c := 10
	return func() int {
		c += 1
		return c
	}
}

//多参函数  能接收任何参数
func args(arg...interface{})  {

	for _,v:= range arg{
		fmt.Println(v)
	}
}

func args2(numbers ...int)  {
	for _,v:= range numbers{
		fmt.Println(v)
	}
}


func main()  {

	f := warps()
	fmt.Println(f())
	fmt.Println(f())

	args("hello",10,time.Now())

	args2([]int{20,30,40,40,100}...)
}
