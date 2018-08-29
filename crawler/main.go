package main

import (
	"imooc.com/learn-go/crawler/engine"
	"imooc.com/learn-go/crawler/scheduler"
	"imooc.com/learn-go/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
