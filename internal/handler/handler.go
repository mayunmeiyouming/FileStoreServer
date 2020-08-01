package handler

import (
	"io"
	"io/ioutil"
	"net/http"
	"log"
)

// UploadHandler ...
func UploadHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("UploadHandler")
	if request.Method == "GET" {
		// 返回上传html页面
		data, err := ioutil.ReadFile("./static/view/upload.html")
		if err != nil {
			log.Println("文件读取失败")
			io.WriteString(response, "internal server error ")
			return
		}
		io.WriteString(response, string(data))
	} else if request.Method == "POST" {
		// 接受文件流及存储到本地目录

	}
}
