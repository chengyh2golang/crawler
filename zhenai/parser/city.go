package parser

import (
	"regexp"
	"crawler/engine"
	"strings"
)

const  cityRe = `<a href="http://m.(zhenai.com/u/[0-9]+)[^d]*" class="left-item"[^>]*><div class="right-item"[^>]*><div class="f-nickname"[^>]*>([^<]+)<span class=`
//const  cityListRe  = `<a href="//m.(zhenai.com/zhenghun/[a-z0-9]+)">([^<]+)</a>`
//http://album.zhenai.com/u/106794356

func ParserCity(contens []byte)  engine.ParserResult {
	re := regexp.MustCompile(cityRe)
	//matches := re.FindAll(contens,-1)

	//-1，找到所有的匹配个数；FindAllSubmatch的返回值[][][]byte 相当于[][]string
	matches := re.FindAllSubmatch(contens, -1)

	result := engine.ParserResult{}
	for _,m := range matches {

		name:= strings.TrimSpace(string(m[2]))

		result.Items = append(result.Items,name)
		//fmt.Println("http://www."+string(m[1]))
		result.Requests = append(result.Requests, engine.Request{
			Url:"http://album." + string(m[1]),
			ParserFunc: func(c []byte) engine.ParserResult {
				return ParseProfile(c,name)
				},
		})

	}
	//fmt.Println(len(matches))
	return result
}
