package scheduler

import (
	engine2 "go_crawler/crawler_concurrent/engine"
)

type QueuedScheduler struct {
	requestChan chan engine2.Request
	workerChan chan chan engine2.Request
}

func (q *QueuedScheduler) WorkerChan() chan engine2.Request {
	return make(chan engine2.Request)
}

func (q *QueuedScheduler) Submit(r engine2.Request) {
	q.requestChan <- r
}

func (q *QueuedScheduler) WorkerReady(w chan engine2.Request) {
	q.workerChan <- w
}

func (q *QueuedScheduler) Run() {
	q.workerChan = make(chan chan engine2.Request)
	q.requestChan = make(chan engine2.Request)
	go func() {
		var requestQ []engine2.Request
		var workerQ []chan engine2.Request
		for {
			var activeRequest engine2.Request
			var activeWorker chan engine2.Request
			if len(requestQ) > 0 && len(workerQ) >0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <- q.requestChan:
				// send r to a ?worker
				requestQ = append(requestQ, r)
			case w := <- q.workerChan:
				// send ?next_request to w
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}