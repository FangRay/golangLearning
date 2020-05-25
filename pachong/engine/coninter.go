package engine

//多线程，带队例，的爬虫

//InterfaceConcurrentegine
type Intercon struct {
	Schedule  Schedule
	WorkerNum int
	ItemChan  chan interface{}
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
		for _, v := range result.Items {
			//log.Printf("result send items value :%v" ,v )
			//这里很关键，要注意v的作用域
			go func(i interface{}) {
				e.ItemChan <- i
			}(v)
		}

		for _, r := range result.Requests {
			if dedup(r.Url) {
				continue
			}
			e.Schedule.Submit(r)

		}

	}

}

func CreatWorker(in chan Requests, out chan ParseResult, rd WorkNotify) {
	go func() {
		for {
			rd.WorkerReady(in)
			r := <-in
			//log.Printf("Workready Mission:%v", r)
			result, err := Worker(r)
			//log.Printf("Workready Result%s", result.Items)
			if err != nil {
				continue
			} else {
				out <- result
			}
		}
	}()

}

//去重复，因为我这里一个地址 会有两个parser，所以不能放在engine里去重
//上面一段因为在gorutine	里对map操作不安全，还是改回来，把true flase值改成数字，取过3次就跳过
var visitedurl = make(map[string]int)

func dedup(url string) bool {
	//log.Println(visitedurl[url])
	if visitedurl[url] == 0 {
		visitedurl[url] = 1
		return false
	}
	if visitedurl[url] == 1 {
		visitedurl[url] = 2
		return false
	}
	if visitedurl[url] == 2 {
		visitedurl[url] = 3
		return false
	}
	if visitedurl[url] == 3 {
		return true
	}
	return false
}
