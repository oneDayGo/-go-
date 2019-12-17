package main

import "fmt"

func defersdemo()  {

	var i int = 1

	//实时计算,也就是直接计算
	defer fmt.Printf("def1=%d",i*2) //i=1
	i++
	defer func() {
		//不实时计算,等到后面执行完参数再计算所以i=3
		fmt.Printf("def2=%d",i*2)
	}()
	i++
}

//返回0 不存在引用关系
func deferReturn() int {
	i := 0

	defer func() {
		i = i+10

	}()
	return i
}

//10 存在引用关系
func deferReturn2()(i int)  {
	i = 0
	defer func() {
		i = 10
	}()
	return 0
}


func main()  {

	//defersdemo()
	fmt.Print(deferReturn())
	fmt.Print(deferReturn2())

}
