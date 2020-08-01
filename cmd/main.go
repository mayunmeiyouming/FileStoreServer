package main
import (
	"log"
	"github.com/mayunmeiyouming/FileStoreServer/internal/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload" ,handler.UploadHandler)
	log.Println("设置监听")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("监听失败: " + err.Error())
	}
} 
