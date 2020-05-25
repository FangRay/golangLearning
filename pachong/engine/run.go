package engine

//单线程版爬虫

import (
	"log"
	"pachong/fetch"
)

type Spengine struct {
}

func (Spengine) Run(seeds ...Requests) {

	var re []Requests
	for _, v := range seeds {
		//fmt.Printf("seeds Request %v",v)
		re = append(re, v)
	}

	for len(re) > 0 {
		r := re[0]
		/*for _, rv :=range re{
			fmt.Printf("现有request%v\n共%d个",rv,len(re))
		}*/
		re = re[1:]
		result, err := Worker(r)
		if err != nil {
			continue
		}
		re = append(re, result.Requests...)
		for _, item := range result.Items {
			log.Printf("got itme %v",
				item)
		}
	}
}

func Worker(r Requests) (ParseResult, error) {
	bytes, err := fetch.Fetch(r.Url)
	if err != nil {
		log.Print(err)
		return ParseResult{}, err
	}
	//log.Printf("Work Fetch Url : %v \n", r.Url)
	return r.ParseFunc(bytes), nil
	//fmt.Printf("Fetch Result\n" ,result.Items)
	// 加。。。表示全部append進去

}
