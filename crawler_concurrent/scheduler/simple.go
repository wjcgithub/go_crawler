package scheduler

import (
	engine2 "go_crawler/crawler_concurrent/engine"
)

type SimpleScheduler struct {
	workerChan chan engine2.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine2.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine2.Request) {

}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine2.Request)
}

func (s *SimpleScheduler) Submit(r engine2.Request) {
	//send request down to worker chan
	go func() {s.workerChan <- r}()
}


