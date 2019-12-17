package main

import "fmt"

/**
实现接口
*/


type Writer interface {
	Write()
}
type Author struct {
	name string
	Writer
}

//定义新结构,重点是实现接口方法Write()
type Other struct {
	i int
}

func (a Author) Write(){
	fmt.Println("author")

}

func (o Other) Write() {
fmt.Println("other")
}


func main() {

	//Other能够赋值给Writer因为他们都共同实现了一个接口Writer
	author := Author{name: "hello", Writer:Other{i: 10}}
	author.Write()
	author.Writer.Write()


}
