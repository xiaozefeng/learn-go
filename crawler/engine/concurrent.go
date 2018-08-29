package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int

}

type Scheduler interface {
	Submit(request Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (c *ConcurrentEngine) Run(seeds ... Request) {

	in := make(chan Request)
	out := make(chan ParseResult)
	c.Scheduler.ConfigureMasterWorkerChan(in)
	for i := 0; i < c.WorkerCount; i++ {
		createWorker(in,out)
	}
	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}
	itemsCount := 0
	for {
		result := <-out
		for _ ,item := range result.Items{
			log.Printf("Got item: %d, %v", itemsCount, item)
			itemsCount ++
		}

		for _, request := range result.Requests {
			c.Scheduler.Submit(request)
		}
	}

}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for{
			request:=<-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
