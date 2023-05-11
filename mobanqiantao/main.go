package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request){
	//定义一个函数，在模板里面使用,要么只有一个返回值，要么第二个返回值是error
	f_moban := func(name string)(string, error){
		return name + "真好看", nil
	}
	//定义模板 f.tmpl
	//解析模板   自定义模板使用自定义函数一定要在解析前定义一下
	//先用template.New创建一个带名字的模板，且创建的模板的名字跟解析的模板名字要一致，然后去解析
	t := template.New("f.tmpl")

	t.Funcs(template.FuncMap{
		"hanshu": f_moban,
	})

	_, err := t.ParseFiles("src/practice/Gin_demo/mobanqiantao/f.tmpl")

	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	name := "zhiqing"
	//渲染模板
	t.Execute(w, name)
}

//嵌套模板
func demo1(w http.ResponseWriter, r *http.Request)  {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("src/practice/Gin_demo/mobanqiantao/t.tmpl", "src/practice/Gin_demo/mobanqiantao/u1.tmpl")
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	name := "zhiqing"
	//渲染模板
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmplDemo", demo1)
	err :=http.ListenAndServe(":9000",nil)
	if err != nil {
		fmt.Printf("http server start failed, err:%v", err)
		return
	}
}
