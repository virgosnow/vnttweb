package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"
)

var tmFileDir = flag.String("d", "", "translation memory file dir")

// transMem 显示翻译记忆
func transMem(w http.ResponseWriter, r *http.Request) {
	display := func(content string) {
		_, err := w.Write([]byte(content))
		if err != nil {
			println(err)
		}
	}
	project := r.URL.Query().Get("project")
	// 无文件直接返回错误提示
	if project == "" {
		display("no specified project")
		return
	}
	projectTmFile := path.Join(*tmFileDir, project+".csv")
	// 读取文件
	bytes, err := os.ReadFile(projectTmFile)
	if err != nil {
		// 出错了的话，能在网页端打印就在网页端打印出来，方便知道问题
		display(fmt.Sprintf("mem file path:[%s] err:[%s]", projectTmFile, err.Error()))
		return
	}
	display(string(bytes))
}

func main() {
	flag.Parse()
	// 超简单http web
	server := http.Server{
		Addr:    ":4655",
		Handler: nil,
	}
	http.HandleFunc("/tm", transMem)

	err := server.ListenAndServe()
	if err != nil {
		println(err)
		return
	}
}
