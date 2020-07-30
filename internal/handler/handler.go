package handler

import (
	"net/http"
	"io/ioutil"
	"ioutil"
	"io"
)

func uploadHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method = "GET" {
		// 返回上传html页面
		ioutil.ReadFile("./static/view/index.html")
	} else if request.Method = "POST" {
		// 接受文件流及存储到本地目录 

	}
}