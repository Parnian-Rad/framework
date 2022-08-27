package ports

type WorkerPool interface {
	RegisterTask(taskName string, handler func(interface{}) interface{}, concurrency int, queueLength int)
	PushJob(taskName string, resultChannel chan interface{}, params interface{})
	Run()
}

type JobRequest struct {
	ResultChannel chan interface{}
	Params        interface{}
}
