package pool

/*
type WorkerPool struct {
	maxWorkers  int
	taskQueue   []chan func()
	stoppedChan chan struct{}
	stopped		bool
	mutex 		sync.Mutex
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
		stopped: 	 false,
	}
	// Start the task dispatcher.
	pool.dispatch()
	pool.dispatchClose()

	return pool
}

func (p *WorkerPool) Submit(uid string, task func()) error {
	p.mutex.Lock()
	if p.stopped {
		p.mutex.Unlock()
		return errors.New("pool closed")
	}
	idx := fnv1a.HashString64(uid) % uint64(p.maxWorkers)
	if task != nil {
		p.taskQueue[idx] <- task
	}
	p.mutex.Unlock()
	return  nil
}

func (p *WorkerPool) dispatch() {
	for i := 0; i < p.maxWorkers; i++ {
		p.taskQueue[i] = make(chan func(), 1024)
		go startWorker(p.taskQueue[i])
	}
}

func (p *WorkerPool) dispatchClose(){
	go func() {
		for !p.stopped {
			<- p.stoppedChan
			p.mutex.Lock()
			for _,v := range p.taskQueue{
				close(v)
			}
			p.stopped = true
			p.mutex.Unlock()
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

func main() {
	pool := New(100)

	go func(pool *WorkerPool) {
		for i:=0; i<100; i++ {
			err := pool.Submit(strconv.Itoa(i), func() {
				fmt.Println(strconv.Itoa(i))
			})
			if err != nil{
				fmt.Println("pool closed")
				return
			}
			time.Sleep(time.Second * 2)
		}
	}(pool)

	time.Sleep(50*time.Second)
	pool.Close()

	time.Sleep(time.Second*10000)
}
 */