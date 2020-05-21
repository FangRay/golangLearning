package engine

//多线程，带队例，的爬虫

import (
	"log"
)

//InterfaceConcurrentegine
type Intercon struct {
	Schedule  Schedule
	WorkerNum int
}

type Schedule interface {
	WorkNotify
	GetWorker() chan Requests
	Submit(Requests)
	Run()
}

type WorkNotify interface {
	WorkerReady(chan Requests)
}

func (e Intercon) Run(seeds ...Requests) {
	e.Schedule.Run()
	out := make(chan ParseResult)
	for i := 0; i < e.WorkerNum; i++ {
		CreatWorker(e.Schedule.GetWorker(), out, e.Schedule)
	}
	for _, r := range seeds {
		e.Schedule.Submit(r)
	}
	for {
		result := <-out
		for _, i := range result.Items {
			log.Printf("ParseValue:%s\n", i)
		}
		for _, r := range result.Requests {
			e.Schedule.Submit(r)
		}
	}

}

func CreatWorker(in chan Requests, out chan ParseResult, rd WorkNotify) {
	go func() {
		for {
			rd.WorkerReady(in)
			r := <-in
			log.Printf("Workready Mission:%v", r)
			result, err := Worker(r)
			log.Printf("Workready Result%s", result.Items)
			if err != nil {
				continue
			} else {
				out <- result
			}
		}
	}()

}
