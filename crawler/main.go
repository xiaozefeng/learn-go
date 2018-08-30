package main

import (
	"imooc.com/learn-go/crawler/engine"
	"imooc.com/learn-go/crawler/scheduler"
	"imooc.com/learn-go/crawler/zhenai/parser"
	"imooc.com/learn-go/crawler/persist"
)

func main() {
	itemSaver, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemSaver,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/guangdong",
	//	ParserFunc: parser.ParseCity,
	//})
}
