package main


/** 接口嵌入接口, 接口不能嵌入结构体**/

type Writer interface {
	Writer()
}
type Reader interface {
	Reader()
}

//接口嵌入接口 (相当于接口继承)
type Teacher interface {
	Writer
	Reader
}



func main() {



}
