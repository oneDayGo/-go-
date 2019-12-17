package main

import (
	"fmt"
	"time"
)

/**
定义通道
channel := make(chan int) 接收又发送的通道
channel := make(chan<-int) 发送通道
channel := make(<-chan int) 接收通道
*/


//写入数据函数
func send(c chan<- int)  {

	for i:=0;i<10;i++{
		fmt.Printf("写入数据%d\n",i)
		//写入数据
		c<-i
		fmt.Printf("写入完成\n")
	}
}

//接收
func recv(c <-chan int)  {

	for i:= range c {
		fmt.Printf("接收到%d\n",i)

	}
}

//无缓冲通道　阻塞
func chan1()  {

	c := make(chan int) //创建无缓冲的通道
	go send(c)
	go recv(c)

	time.Sleep(3*time.Second)
	close(c) //关闭通道

}

//有缓冲通道 不阻塞
func chan2()  {

	c := make(chan int,10)
	go send(c)
	go recv(c)
	time.Sleep(3*time.Second)
	close(c)


}


func main() {
	//无缓冲通道
	//chan1()

	chan2()
}
