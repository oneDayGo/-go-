package main

import "fmt"

type MyType int

func (i MyType) prints() {
	fmt.Printf("自定义类型方法%d",i)
}

func main() {

	var i MyType
	i.prints()

	number := MyType(10)
	number.prints()


}
