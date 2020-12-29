package parser

import (
	"regexp"

	"github.com/lingjiao0710/test/crawler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

func PrintCityList(contents []byte) engine.ParseResult {
	compile := regexp.MustCompile(cityListRe)
	matchs := compile.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matchs {
		//fmt.Printf("m[0] : %s\n", m[0])
		//fmt.Printf("city: %s, url %s\n", m[2], m[1])

		//城市名加到items内
		result.Items = append(result.Items, string(m[2]))
		//城市URL
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}

	return result
	//fmt.Printf("count %d\n", len(matchs))
}
