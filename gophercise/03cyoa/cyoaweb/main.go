package main

import (
	"flag"
	"fmt"
	cyoa "github.com/lingjiao0710/test/gophercise/03cyoa"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "WEB服务端口号")
	fileName := flag.String("file", "gopher.json", "Choose your own Adventure故事的json文件")
	flag.Parse()

	fmt.Printf("Using the stroy in %s\n", *fileName)

	f, err := os.Open(*fileName)
	if err != nil {
		fmt.Printf("打开%s失败！\n", *fileName)
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		fmt.Printf("JsonStory%s失败！\n", *fileName)
		panic(err)
	}

	//fmt.Printf("story :%+v\n", story)

	//默认使用story的tmpl
	h := cyoa.NewHandler(story)

	//指定tmpl

	//tmpl := template.Must(template.New("").Parse("你好啊！"))
	//h := cyoa.NewHandler(story, cyoa.WithTemplate(tmpl))
	fmt.Printf("服务启动在端口：%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
