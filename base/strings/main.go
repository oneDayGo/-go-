package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)
//字符串长度
func str()  {
	s := "hellowrold你好世界"

	//获取字节长度
	fmt.Println(len(s))
	//获取字符长度
	fmt.Print(utf8.RuneCountInString(s))

	//循环的是每一个rune
	for k,v := range s{
		//k %d是索引 v %d是unicode码值
		fmt.Printf("\n%d===%c==%d\n",k,v,v)
	}
}

//拼接字符串
func joins()  {
	fmt.Printf("hello"+"wrold\n") //性能比较差
	fmt.Printf(fmt.Sprintf("%d:%s",2019,"年\n") )//性能一般
	fmt.Printf(strings.Join([]string{"hello","world\n"},","))  //以逗号拼接,效率高,但构建一个本来没有的数据代价也不小

	//比较理想的方式
	var buffer bytes.Buffer
	buffer.WriteString("hello")
	buffer.WriteString(",")
	buffer.WriteString("wrold")
	fmt.Printf(buffer.String())

	//如果可以预估长度的话,推荐使用这种
	var bl strings.Builder
	bl.WriteString("ABC")
	bl.WriteString("CDB")
	fmt.Printf(bl.String())

}

//strings包
func stringsLib()  {

	s := "helloworld"

	/*fmt.Println(strings.HasPrefix(s,"h")) //判断字符串是否以某个字符开头
	fmt.Println(strings.HasSuffix(s,"l")) //以某个字符为结尾
	fmt.Println(strings.Split(s,"o")) //分割字符串
	fmt.Println(strings.Index(s,"o")) //返回字符串索引
	fmt.Println(strings.LastIndex(s,"o")) //最后一个匹配索引

	 */

	str := make([]string,0)
	for _,v := range []rune(s){
		str = append(str,string(v))
	}
	fmt.Println(strings.Join(str,",,,")) //字符串连接

	fmt.Println(strings.Replace(s,"o","k",-1)) //n小于１会替换所有
	fmt.Println(strings.Count(s,"o")) //统计字出现的个数
	fmt.Println(strings.Contains(s,"o")) //判断字符包含


}

func main()  {
	//joins()
	stringsLib()
}
