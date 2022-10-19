package goqueue

type ConcurrentQueue struct {
	queue chan string 
}

func NewQueue() *ConcurrentQueue {
	return &ConcurrentQueue{
	}
}

func (Q *ConcurrentQueue) Init(capacity int) () {
	Q.queue = make(chan string, capacity)
}

func (Q *ConcurrentQueue) Enqueue(item string) (int) {
	Q.queue<-item
	return 0
}

func (Q *ConcurrentQueue) Dequeue() (string, int) {
	var ok bool
	var res string
	var err int = 0
	res, ok = <-Q.queue
	if !ok {
		err = 1
	}
	return res, err
}

func (Q *ConcurrentQueue) Size() (int) {
	return len(Q.queue)
}

func (Q *ConcurrentQueue) Capacity() (int) {
	return cap(Q.queue)
}
