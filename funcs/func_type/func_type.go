package main

import (
	"fmt"
	"time"
)

//定义一个类型
type funcType func(time.Time)

func CureeTime( start time.Time)  {

	fmt.Println(start)
}

func main()  {
	//先把CureeTime函数转为funcType类型,然后传入参数
	funcType(CureeTime)(time.Now())
}
