package main

import (
	"fmt"
	"reflect"
)

/**
获取结构体的tag

*/

type Tags struct {
	name string "姓名"
	age int "年龄"
	sex string "性别"
}

func main() {

	//通过反射获取tag
	var tag Tags
	fmt.Println(reflect.TypeOf(tag).Field(0).Tag)
	fmt.Println(reflect.TypeOf(tag).Field(1).Tag)
	fmt.Println(reflect.TypeOf(tag).Field(2).Tag)
}


