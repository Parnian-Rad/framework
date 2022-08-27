package asynq

import (
	"encoding/json"
	"time"

	"git.snapp.ninja/search-and-discovery/framework/pkg/ports"
	"github.com/hibiken/asynq"
)

type distributedWorkerpoolClient struct {
	redisAddress string
	client       *asynq.Client
}

func New(redisAddress string) ports.DistributedWorkerpoolClient {
	return &distributedWorkerpoolClient{
		redisAddress: redisAddress,
		client:       asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddress}),
	}
}

func (dwc *distributedWorkerpoolClient) PushJob(taskName string, params interface{}, maxRetry int, timeOut time.Duration) error {
	var paramsB []byte
	var err error
	if params != nil {
		paramsB, err = json.Marshal(params)
		if err != nil {
			return err
		}
	}
	_, err = dwc.client.Enqueue(asynq.NewTask(taskName, paramsB))
	return err
}
