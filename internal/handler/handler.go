package handler

import (
	"os"
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
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			log.Println("文件读取失败")
			io.WriteString(response, "internal server error ")
			return
		}
		io.WriteString(response, string(data))
	} else if request.Method == "POST" {
		// 接受文件流及存储到本地目录
		log.Println("文件上传中")
		uploadFile, uploadFileHeader, err := request.FormFile("file")
		if err != nil {
			log.Println("文件上传出错")
			response.WriteHeader(500)
			io.WriteString(response, "internal server error ")
			return
		}
		defer uploadFile.Close()

		newFile, err := os.Create("/tmp/" + uploadFileHeader.Filename)
		if err != nil {
			log.Println("新文件创建失败")
			response.WriteHeader(500)
			io.WriteString(response, "internal server error ")
			return
		}
		defer newFile.Close()
		_, err = io.Copy(newFile, uploadFile)
		if err != nil {
			log.Println("文件保存失败")
			response.WriteHeader(500)
			io.WriteString(response, "internal server error ")
			return
		}
		http.Redirect(response, request, "/file/upload/suc", http.StatusFound)
		log.Println("文件上传成功")
		return
	}
	log.Println("未知错误: ", request.Method)
}

// UploadSucHandler ...
func UploadSucHandler(response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, "Upload Finished!")
}
