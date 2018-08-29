package parser

import (
	"imooc.com/learn-go/crawler/engine"
	"regexp"
)

const cityReg = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityReg)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User:"+name)
		result.Requests  = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
				// 利用闭包的概念，用函数包装一下
				return ParseProfile(bytes, name)
			},
		})
	}

	return result
}
