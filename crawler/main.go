package main

import (
	"imooc.com/learn-go/crawler/engine"
	"imooc.com/learn-go/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})


}
