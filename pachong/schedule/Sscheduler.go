package schedule

import (
	"pachong/engine"
)

type Sscheduler struct {
	RequestChan chan engine.Requests
	WorkerChan  chan chan engine.Requests
}

func (s *Sscheduler) GetWorker() chan engine.Requests {
	return make(chan engine.Requests)
}

func (s *Sscheduler) Submit(c engine.Requests) {
	s.RequestChan <- c
}

func (s *Sscheduler) WorkerReady(w chan engine.Requests) {
	s.WorkerChan <- w
}

func (s *Sscheduler) Run() {
	s.WorkerChan = make(chan chan engine.Requests)
	s.RequestChan = make(chan engine.Requests)

	var requestQ []engine.Requests
	var workerQ []chan engine.Requests

	go func() {
		for {
			//这两句要写在for里面，当nil的时候，不会被select到
			var activerequest engine.Requests
			var activeworker chan engine.Requests

			if len(requestQ) > 0 && len(workerQ) > 0 {
				activerequest = requestQ[0]
				activeworker = workerQ[0]
			}
			select {
			case r := <-s.RequestChan:
				requestQ = append(requestQ, r)
			case w := <-s.WorkerChan:
				workerQ = append(workerQ, w)
			case activeworker <- activerequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]

			}
		}
	}()
}
