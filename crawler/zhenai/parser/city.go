package parser

import (
	"imooc.com/learn-go/crawler/engine"
	"regexp"
)

var (
	profileReg = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlReg = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {
	matches := profileReg.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		//result.Items = append(result.Items, "User:"+name)
		url := string(m[1])
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(bytes []byte) engine.ParseResult {
				// 利用闭包的概念，用函数包装一下
				return ParseProfile(bytes, url,  name)
			},
		})
	}

	cityMatches := cityUrlReg.FindAllSubmatch(contents, -1)
	for _, m := range cityMatches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
