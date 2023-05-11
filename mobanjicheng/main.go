package main

import (
	"fmt"
	"html/template"
	"net/http"
)
//block 实现模板继承

func index(w http.ResponseWriter, r *http.Request){
	//定义模板
	//解析模板
	t, err := template.ParseFiles("src/practice/Gin_demo/mobanjicheng/index.tmpl")
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	//渲染模板
	msg := "这是index页面"
	t.Execute(w, msg)
}
func home(w http.ResponseWriter, r *http.Request){
	//定义模板
	//解析模板
	t, err := template.ParseFiles("src/practice/Gin_demo/mobanjicheng/home.tmpl")
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	//渲染模板
	msg := "这是home页面"
	t.Execute(w, msg)
}

func index_base(w http.ResponseWriter, r *http.Request){
	//定义模板(使用模板继承的方式)
	//解析模板  按照模板包含与被包含的关系，把需要解析的模板名字都写上
	t, err := template.ParseFiles("src/practice/Gin_demo/mobanjicheng/base.tmpl", "src/practice/Gin_demo/mobanjicheng/index_base.tmpl")
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	//渲染模板
	msg := "这是index_baseh2页面"
	//t.Execute(w, msg)可以直接使用这个显示，且正常显示
	t.ExecuteTemplate(w, "index_base.tmpl", msg)
}
func home_base(w http.ResponseWriter, r *http.Request){
	//定义模板(使用模板继承的方式)
	//解析模板
	t, err := template.ParseFiles("src/practice/Gin_demo/mobanjicheng/base.tmpl", "src/practice/Gin_demo/mobanjicheng/home_base.tmpl")
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	//渲染模板
	msg := "这是home_baseh2页面"
	//t.Execute(w, msg)可以直接使用这个显示，且正常显示
	t.ExecuteTemplate(w, "home_base.tmpl", msg)
}
func base(w http.ResponseWriter, r *http.Request){
	//定义模板(使用模板继承的方式)
	//解析模板
	t, err := template.ParseFiles("src/practice/Gin_demo/mobanjicheng/base.tmpl")
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	//渲染模板
	msg := "这是base页面"
	t.Execute(w, msg)
}

func main() {
	http.HandleFunc("/index", index)//home和index页面的布局相同的，所以可以采用模板继承的方式进行
	http.HandleFunc("/home", home)
	http.HandleFunc("/index_base", index_base)
	http.HandleFunc("/home_base", home_base)
	http.HandleFunc("/base", base)
	err :=http.ListenAndServe(":9000",nil)
	if err != nil {
		fmt.Printf("http server start failed, err:%v", err)
		return
	}
}
