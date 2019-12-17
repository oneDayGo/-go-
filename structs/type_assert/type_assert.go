package main

import "fmt"

/** 类型断言*/

func assets(i interface{})  {
	//类型断言
	switch str:=i.(type) {

	case string:
		fmt.Printf("%s==为字符串\n",str)
	case int:
		fmt.Printf("%d==为int类型\n",str)
	default:
		fmt.Println("不知道什么类型")
	}
}

func main()  {

	assets(10)
	assets("hello")

	//第二种类型断言
	var str interface{}
	str = "hello"

	if s,ok:= str.(string);ok{
		fmt.Println("是字符串类型")
		fmt.Println(s)
	}

	//精类型断言
	str2:= "hello"
	if _,ok := interface{}(str2).(string);ok{
		fmt.Printf("是字符串类型")
	}


}
