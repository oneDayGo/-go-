package main

import (
	"fmt"
	"reflect"
)

/**
方法提升
*/

type People struct {
	Age int
	Name string
	gender string
}

type OtherPeople struct {
	People
}

//必须是大写才能提成
func (p People) Pe() {
	fmt.Println("pe")
}
func (p *People)Pe2()  {
	fmt.Println("pe2")
}


type NewPeople People


func (p *NewPeople)PeName() {
	fmt.Println("name")
}

func promoteMethors(a interface{})  {
	t := reflect.TypeOf(a)
	for i,n:=0,t.NumMethod();i<n;i++{
		fmt.Println(t.Method(i).Type,t.Method(i).Name)
	}
}



func main()  {


	other :=  OtherPeople{People{Name:"jede",Age:100,gender:"男"}}
    nPeole :=NewPeople{}

    promoteMethors(&other) //可以调用两个方法,值类型和指针类型
	promoteMethors(other) //只能调用值类型方法
	promoteMethors(nPeole)








}
