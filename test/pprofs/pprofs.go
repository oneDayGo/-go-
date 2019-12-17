package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

//监控堆栈cpu等信息

func myfunc(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"hi")
}
func main()  {

	http.HandleFunc("/",myfunc)
	http.ListenAndServe(":8080",nil)
}
