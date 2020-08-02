package main
import (
	"log"
	"github.com/mayunmeiyouming/FileStoreServer/internal/handler"
	"github.com/mayunmeiyouming/FileStoreServer/internal/config"
	"net/http"
)

func main() {
	config := config.New()
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(config.RootPath + "/static"))))
	log.Println("设置监听")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("监听失败: " + err.Error())
	}
} 
