package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(request Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (c *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	c.Scheduler.Run()
	for i := 0; i < c.WorkerCount; i++ {
		createWorker(c.Scheduler.WorkerChan(),
			out,
			c.Scheduler)

	}
	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}
	for {
		result := <-out
		for _, item := range result.Items {
			go func() {
				c.ItemChan <- item
			}()
		}

		for _, request := range result.Requests {
			// 判断是否重复
			c.Scheduler.Submit(request)
		}
	}

}

func createWorker(in chan Request,
	out chan ParseResult,
	notifier ReadyNotifier) {

	go func() {
		for {
			notifier.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
