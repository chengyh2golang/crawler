package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

func main() {

	//"http://m.zhenai.com/zhenghun"

	engine.Run(engine.Request{
		Url:"http://m.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})

}
