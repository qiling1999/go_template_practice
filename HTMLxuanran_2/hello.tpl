<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
    <p>u1</p>
    <p>姓名 {{.u1.Name}}</p>
    <p>性别 {{.u1.Gender}}</p>
    <p>年龄 {{.u1.Age}}</p>
    {{/*这是注释，
    可以换行*/}}
    <p>m1</p>
    <p>姓名 {{.m1.name}}</p>
    <p>性别 {{.m1.gender}}</p>
    <p>年龄 {{.m1.age}}</p>
<hr>{{/*换行符*/}}
{{ $v1 := 100}}{{/*定义一个变量*/}}
{{ $age := .m1.age}}{{/*定义一个变量把前面的变量接收起来*/}}

<hr>
{{ if $v1 }}
{{ $v1 }}
{{ else }}
    啥都没有
{{ end }}
    <hr>
{{ if lt .m1.age 22 }}
好好上学
{{ else }}
好好工作
{{ end }}
<hr>
{{range $idx, $hobby := .hobby}}
    <p>{{$idx}} - {{$hobby}}</p>
{{else}}
    没啥爱好
{{end}}

<hr>
    <p>u1</p>
    {{with .u1}}
    <p>姓名 {{.Name}}</p>
    <p>性别 {{.Gender}}</p>
    <p>年龄 {{.Age}}</p>
{{end}}
</body>
</html>