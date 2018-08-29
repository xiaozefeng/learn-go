package engine

import (
	"fmt"
	"imooc.com/learn-go/crawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		// 将slice中的内容加入requests
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			fmt.Printf("Got itme: %v \n", item)
		}

	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching: %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error, fetching url %s, %v", r.Url, err)
		return ParseResult{}, err
	}
	parseResult := r.ParserFunc(body)
	return parseResult, nil
}
