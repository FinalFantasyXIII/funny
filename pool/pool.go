package pool

import (
	"github.com/segmentio/fasthash/fnv1a"
)

type WorkerPool struct {
	maxWorkers  int
	taskQueue   []chan func()
	stoppedChan chan struct{}
}

func New(maxWorkers int) *WorkerPool {
	// There must be at least one worker.
	if maxWorkers < 1 {
		maxWorkers = 1
	}

	// taskQueue is unbuffered since items are always removed immediately.
	pool := &WorkerPool{
		taskQueue:   make([]chan func(), maxWorkers),
		maxWorkers:  maxWorkers,
		stoppedChan: make(chan struct{}),
	}
	// Start the task dispatcher.
	pool.dispatch()
	pool.dispatchClose()

	return pool
}

func (p *WorkerPool) Submit(uid string, task func()) {
	idx := fnv1a.HashString64(uid) % uint64(p.maxWorkers)
	if task != nil {
		p.taskQueue[idx] <- task
	}
}

func (p *WorkerPool) dispatch() {
	for i := 0; i < p.maxWorkers; i++ {
		p.taskQueue[i] = make(chan func(), 1024)
		go startWorker(p.taskQueue[i])
	}
}

func (p *WorkerPool) dispatchClose(){
	go func() {
		for {
			<- p.stoppedChan
			for _,v := range p.taskQueue{
				close(v)
			}
			p.maxWorkers = 0
			p.stoppedChan <- struct{}{}
		}
	}()
}

func (p *WorkerPool) Close(){
	p.stoppedChan <- struct{}{}
	<- p.stoppedChan
	close(p.stoppedChan)
}

func startWorker(taskChan chan func()) {
	for f := range taskChan {
		f()
	}
}
