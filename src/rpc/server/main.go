package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	// http://127.0.0.1:8080/add?a=1&b=2
	// 返回格式化 json {}
	// 1、call id 的问题  r.Url.Path
	// 2、数据的传输协议 url参数传输协议 http的协议【超文本传输协议】底层tcp协议
	// 3、网路传输协议 http
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()               //解析参数
		fmt.Println("path", r.URL.Path) // call_id
		a, _ := strconv.Atoi(r.Form["a"][0])
		b, _ := strconv.Atoi(r.Form["b"][0])
		w.Header().Set("Content-Type", "application/json")
		jData, _ := json.Marshal(map[string]int{
			"data": a + b,
		})
		_, _ = w.Write(jData)
		_ = http.ListenAndServe(":7000", nil)
	})
}
