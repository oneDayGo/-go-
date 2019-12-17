package main

import (
	"errors"
	"fmt"
)

/**
错误
*/
//第一种返回错误方式
func throwErros() (int,error) {
	return 10,errors.New("出现错误")

}
//第二种
func throwErrorf()(int,error) {

	return  1,fmt.Errorf("返回错误")
}

func main()  {

	if _,err := throwErros();err!=nil{
		fmt.Println(err)
	}
	if _,err:= throwErrorf();err!=nil{
		fmt.Println(err)
	}


}
