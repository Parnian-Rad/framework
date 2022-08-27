package workerpool

import "git.snapp.ninja/search-and-discovery/framework/pkg/ports"

type NativeWorkerpool struct {
	log   ports.Logger
	tasks map[string]*task
}

type task struct {
	log          ports.Logger
	name         string
	concurrency  int
	queueLength  int
	requestQueue chan ports.JobRequest
	handler      func(interface{}) interface{}
}

func New(log ports.Logger) ports.WorkerPool {
	return &NativeWorkerpool{
		log:   log,
		tasks: map[string]*task{},
	}
}

func (t *task) worker() {
	t.log.Info("Worker %s initialized", t.name)
	for job := range t.requestQueue {
		job.ResultChannel <- t.handler(job.Params)
	}
	t.log.Warn("Worker %s initialized", t.name)
}

func (nw *NativeWorkerpool) RegisterTask(taskName string, handler func(interface{}) interface{}, concurrency int, queueLength int) {
	nw.tasks[taskName] = &task{
		log:          nw.log,
		name:         taskName,
		concurrency:  concurrency,
		queueLength:  queueLength,
		requestQueue: make(chan ports.JobRequest, concurrency),
		handler:      handler,
	}
	nw.log.Info("task %s registered", taskName)
}

func (nw *NativeWorkerpool) PushJob(taskName string, resultChannel chan interface{}, params interface{}) {
	nw.tasks[taskName].requestQueue <- ports.JobRequest{
		ResultChannel: resultChannel,
		Params:        params,
	}
	nw.log.Info("job %s pushed", taskName)
}

func (nw *NativeWorkerpool) Run() {
	for _, task := range nw.tasks {
		for i := 0; i < task.concurrency; i++ {
			go task.worker()
		}
	}
}
