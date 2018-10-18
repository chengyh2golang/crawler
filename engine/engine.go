package engine

import (

	fetcher2 "crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {

	//创建一个队列
	var requests []Request

	//将seed存入队列
	for _,seed := range seeds {

		requests = append(requests,seed)
	}

	//队列中有任务就做事情
	for len(requests) > 0 {

		//开始取队列中的第一个元素
		r := requests[0]
		requests = requests[1:]

		//打印这个任务要爬取的url
		log.Printf("fetching url : %s",r.Url)

		//将url传入fetch中获得网页的body
		body, err := fetcher2.Fetcher(r.Url)


		//如果获取当前网页body有错，打印信息后继续下一次循环
		if err != nil {
			log.Printf("fetch err from :%s %v",r.Url,err)
			continue
		}

		//将拿到的body传入这个request的ParserFunc
		parseResult := r.ParserFunc(body)

		//将ParserFunc返回的request信息再加入到任务队列中
		requests = append(requests,parseResult.Requests...)

		//将ParserFunc返回的item信息打印出来
		for _,item := range parseResult.Items {
			log.Printf("got item %v",item)
		}

	}



}