package main

import "fmt"

func main()  {

	//定义一个数组
	x:=[3]int{1,2,4}

	func(arr *[3]int){
		//改变索引0的值
		(*arr)[0] = 10
	}(&x)

	fmt.Println(x)
}
