package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
	"crawler/scheduler"
)

func main() {

	//"http://m.zhenai.com/zhenghun"

	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:"http://m.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParserCityList,
	//})

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 10,

	}
	e.Run(engine.Request{
		Url:"http://m.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})

}
