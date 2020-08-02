package main

import (
	"log"
	"net/http"

	"github.com/mayunmeiyouming/FileStoreServer/internal/config"
	"github.com/mayunmeiyouming/FileStoreServer/internal/handler"
)

func handleListenPath(config config.Config) {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(config.RootPath+"/static"))))
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
}

func main() {
	config := config.New()
	handleListenPath(config)
	log.Println("设置监听")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("监听失败: " + err.Error())
	}
}
