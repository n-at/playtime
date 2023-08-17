package heartbeatpool

type Task func()

type Pool struct {
	size  int
	tasks chan Task
	stop  chan bool
}

func New(threads int) *Pool {
	pool := &Pool{
		size:  threads,
		tasks: make(chan Task),
		stop:  make(chan bool),
	}

	for i := 0; i < threads; i++ {
		go func() {
			for {
				select {
				case task := <-pool.tasks:
					task()
				case <-pool.stop:
					return
				}
			}
		}()
	}

	return pool
}

func (p *Pool) Stop() {
	for i := 0; i < p.size; i++ {
		p.stop <- true
	}
}

func (p *Pool) Add(task Task) {
	p.tasks <- task
}
