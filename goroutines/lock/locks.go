package main

import (
	"fmt"
	"sync"
	"time"
)

/** 锁*/

//互斥锁
func mutexLock()  {
	//用于线程总同步
	wg := sync.WaitGroup{}
	//互斥锁
	var mutex sync.Mutex
	for i:=1;i<4;i++{
		wg.Add(1)
		go func(i int) {
			mutex.Lock()
			fmt.Printf("锁住%d",i)
			time.Sleep(1000)
			mutex.Unlock()
			fmt.Printf("释放锁%d",i)
			wg.Done()
		}(i)
	}

	wg.Wait() //等待执行完毕
	fmt.Printf("执行完毕")
}


//读写锁
func reLock()  {

	m := sync.RWMutex{} //读写锁
	wg := sync.WaitGroup{}
	Data := 0

	wg.Add(20)
	for i:=0;i<10;i++{

		go func(t int) {
			//读锁
			m.RLock()
			defer m.RUnlock()
			fmt.Printf("读数据%d",Data)
			wg.Done()
			time.Sleep(time.Second)
		}(i)

		go func(t int) {

			m.Lock()
			defer m.Unlock()
			Data += t
			fmt.Printf("写数据%d",Data)
			wg.Done()
			time.Sleep(time.Second)
		}(i)
	}

	wg.Wait()
	fmt.Printf("执行完毕")
}

//只执行一次锁
func oneLock()  {

	var once sync.Once
	for i:=0;i<10;i++{
		go func(i int) {
			//只执行一次
			once.Do(func() {
				fmt.Printf("执行了%d",i)
			})
		}(i)
	}
	time.Sleep(time.Second)
}

//map锁 是线程安全
func mapLock()  {

	var m sync.Map
	//保存数据
	m.Store("name","joe")
	m.Store("gender","Male")

	//若key不存在则写入,返回false和输入vlaue
	v,ok := m.LoadOrStore("name1","jin")
	v1,ok2:= m.LoadOrStore("name","jo")
	fmt.Printf("value===%s====fig===%s\n",v,ok)
	fmt.Printf("value===%s=====fig====%s\n",v1,ok2)

	//读取数据
	if v,ok:= m.Load("name");ok{
		fmt.Printf("%s\n",v)
	}

	//遍历所有数据
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("key====%s====value===%s\n",key,value)
		return  true
	})

	//删除数据
	//m.Delete("name1)

}


func main() {
	//mutexLock()
	//reLock()
	//oneLock()
	mapLock()

}
