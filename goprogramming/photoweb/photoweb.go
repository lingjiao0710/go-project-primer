package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	UPLOAD_DIR = "uploads"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	/*结合main()和uploadHandler()方法，针对 HTTP GET 方式请求 /upload 路径，
	程序将会往http.ResponseWriter类型的实例对象w中写入一段HTML文本，即输出一个HTML
	上传表单。如果我们使用浏览器访问这个地址，那么网页上将会是一个可以上传文件的表单。*/
	if r.Method == "GET" {
		io.WriteString(w, "<form method=\"POST\" action=\"/upload\" "+
			" enctype=\"multipart/form-data\">"+
			"Choose an image to upload: <input name=\"image\" type=\"file\" />"+
			"<input type=\"submit\" value=\"Upload\" />"+
			"</form>")
		return
	}

	/*如果是客户端发起的HTTP POST 请求，那么首先从表单提交过来的字段寻找名为 image 的文
	件域并对其接值，调用r.FormFile()方法会返回3个值，各个值的类型分别是multipart.File、
	*multipart.FileHeader和error。
	如果上传的图片接收不成功，那么在示例程序中返回一个HTTP服务端的内部错误给客户端。
	如果上传的图片接收成功，则将该图片的内容复制到一个临时文件里。如果临时文件创建失
	败，或者图片副本保存失败，都将触发服务端内部错误。
	如果临时文件创建成功并且图片副本保存成功，即表示图片上传成功，就跳转到查看图片页
	面。此外，我们还定义了两个defer语句，无论图片上传成功还是失败，当uploadHandler()
	方法执行结束时，都会先关闭临时文件句柄，继而关闭图片上传到服务器文件流的句柄。
	别忘了在程序开头引入io/ioutil这个包，因为示例程序中用到了ioutil.TempFile()这
	个方法。*/

	/*
		postman 使用post方式上传图片
		Headers 设置Content-Type = "multipart/form-data"
		Body 设置form-data，key = "image",value选择本地图片

	*/

	if r.Method == "POST" {
		file, header, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fileName := header.Filename
		defer file.Close()

		//使用filepath.join可以生成对应操作系统的路径
		createFile := filepath.Join(UPLOAD_DIR, fileName)
		fmt.Println(createFile)
		t, err := os.Create(createFile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer t.Close()
		if _, err := io.Copy(t, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "view?id="+fileName, http.StatusFound)
	}

}

//浏览器使用get请求，获取id为图片文件名：
//http://localhost:8080/view?id=20141119212339_seHfw.jpeg
func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := filepath.Join(UPLOAD_DIR, imageId)
	fmt.Println(imagePath)

	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	fmt.Println(os.IsExist(err))
	if err == nil {
		return true
	}

	return os.IsExist(err)
}

func main() {

	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/view", viewHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
