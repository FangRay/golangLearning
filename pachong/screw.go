package main

//
import (
	"pachong/biquge/parse"
	"pachong/engine"
	"pachong/schedule"
)

func main() {

	ce := engine.Intercon{
		WorkerNum: 100,
		Schedule:  &schedule.Sscheduler{},
	}

	// 这个是所有work都用一个in的chan
	/*ce := engine.Intercon{
		WorkerNum: 10,
		Schedule:  &schedule.SimpleSchedule{},
	}*/

	//ce:=engine.ConcurrentEgine{}
	//单线程
	//ce:=engine.Spengine{}
	ce.Run(engine.Requests{
		Url:       "https://www.biquge.com.cn",
		ParseFunc: parse.ParseMain,
	})

}
