package main

import (
	"github.com/lingjiao0710/test/crawler/engine"
	"github.com/lingjiao0710/test/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.PrintCityList,
	})

	/*resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	//body在退出时需要关闭
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error, status code is ", resp.StatusCode)
		return
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", all)
	printCityList(all)*/
}

func printCityList(contents []byte) {
	/*exp := `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`
	compile := regexp.MustCompile(exp)
	matchs := compile.FindAllSubmatch(contents, -1)

	for _, m := range matchs {
		//fmt.Printf("m[0] : %s\n", m[0])
		fmt.Printf("city: %s, url %s\n", m[2], m[1])
	}

	fmt.Printf("count %d\n", len(matchs))*/
}
