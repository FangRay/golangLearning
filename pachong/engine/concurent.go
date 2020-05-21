package engine

//我自己写的多纯种爬虫
import (
	"log"
)

type ConcurrentEgine struct {
}

func (c ConcurrentEgine) Run(seeds ...Requests) {

	in := make(chan Requests)
	out := make(chan ParseResult)

	for _, r := range seeds {
		go func() {
			in <- r
		}()
	}

	for i := 0; i < 50; i++ {
		go func() {
			for {
				re := <-in
				//log.Printf("request from in, Url: %s, Func: %v\n ",re.Url,re.ParseFunc)
				result, err := Worker(re)
				if err != nil {
					continue
				} else {
					out <- result
				}
			}
		}()
	}

	for {
		result := <-out
		for _, r := range result.Items {
			log.Printf("Get Result:%s /n", r)
		}
		for _, q := range result.Requests {
			// 好大一个坑，for里面传参，一定，一定，一定要复制，不然就死B
			request := q
			go func() {
				in <- request
			}()
		}

	}

}
