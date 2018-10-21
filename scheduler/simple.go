package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	WorkChan chan engine.Request
}

func (s *SimpleScheduler) ConfigMasterWorkerChan(c chan engine.Request) {
	s.WorkChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {s.WorkChan <- r}()
}



