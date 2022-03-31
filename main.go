package main

import (
	"fmt"      // fmt 包含有格式化 I/O 函数。主要分为向外输出内容和获取输入内容两大部分
	"net/http" // 提供了 HTTP 编程有关的接口，内部封装了 TCP 连接和报文解析的复杂琐碎细节。http 提供了 HTTP 客户端和服务器实现。
)

// http.Request是用户的请求信息，一般用 r 作为简写
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint 将 <h1>Hello, 这里是 goblog</h1> 子串写入 http.ResponseWriter，即可响应用户请求
	//fmt.Fprint(w, "<h1>Hello, 这里是 goblog</h1>")
	//fmt.Fprint(w, "请求路径为："+r.URL.Path)

	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello,这里是Ruan goblog</h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprint(w, "此博客用于记录阿翔的编程笔记，如果有什么建议，请联系 "+"<a href=\"mailto:734162396@qq.com\">RuanZzzz@example.com</a>")
	} else {
		fmt.Fprint(w, "<h1>未找到请求页面 :(</h1>"+"<p>如果有疑惑请联系我们</p>")
	}

}
func main() {
	// 用以指定处理 HTTP 请求的函数
	// http.HandleFunc 里传参的 / 表示：任意路径
	http.HandleFunc("/", handlerFunc)
	// 监听本地 3000 端口以提供服务
	http.ListenAndServe(":3000", nil)
}
