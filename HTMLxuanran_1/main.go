package main
//Gin渲染
//首先学习一下go语言原生标准库http/template 模板与渲染  模板引擎
//模板可以理解为，事先定义好一个HTML文档，，模板渲染的作用机制可以简单理解为
//文本替换操作-使用相应的数据去替换HTML文档中事先准备好的标记
//1.模板文件后缀为.tmpl和.tpl为后缀，必须使用UTF8编码2.使用{{和}}包裹和标识需要传入的数据，通.访问传给模板的数据
//模板文件使用分为三个部分：1定义，2解析，3渲染
//解析函数 Parse() ParseFiles()  ParseGlob()
//渲染模板 Execute() ExecuteTemplate()

import (
	"fmt"
	"html/template"
	"net/http"
)


//先用原生的框架做一个模板渲染，后面再用gin框架做模板渲染
func sayHello(w http.ResponseWriter, r *http.Request){
	//模板已经定义好了，就是hello.tpl或者hello.html
	//2解析模板
	t, err :=template.ParseFiles("src/practice/Gin_demo/HTMLxuanran_1/hello.tpl")
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	//3渲染模板
	err = t.Execute(w, "helloTest") //这里传的是一个字符串
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
}
func main() {
	http.HandleFunc("/",sayHello)
	err :=http.ListenAndServe(":9000",nil)
	if err != nil {
		fmt.Printf("http server start failed, err:%v", err)
		return
	}
}
