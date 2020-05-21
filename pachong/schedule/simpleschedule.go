package schedule

import (
	"log"
	"pachong/engine"
)

type SimpleSchedule struct {
	WorkChan chan engine.Requests
}

func (s *SimpleSchedule) WorkerReady(chan engine.Requests) {
	log.Printf("我是简单的schedule 共用的一个schedule线程")
}

func (s *SimpleSchedule) GetWorker() chan engine.Requests {
	return s.WorkChan
}

func (s *SimpleSchedule) Submit(r engine.Requests) {
	//这里送任务的时候，要单独开一个线程，要不就会阻塞
	go func() {
		s.WorkChan <- r
	}()
}

func (s *SimpleSchedule) Run() {
	s.WorkChan = make(chan engine.Requests)
}
