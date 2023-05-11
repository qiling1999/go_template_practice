package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/base", base)
	err :=http.ListenAndServe(":9000",nil)
	if err != nil {
		fmt.Printf("http server start failed, err:%v", err)
		return
	}
}