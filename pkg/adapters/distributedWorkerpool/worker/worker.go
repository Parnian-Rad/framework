package asynq

import (
	"context"
	"net/http"

	"git.snapp.ninja/search-and-discovery/framework/pkg/ports"
	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
)

type distributedWorkerpoolWorker struct {
	redisAddress string
	server       *asynq.Server
	router       *asynq.ServeMux
}

func New(redisAddress string, concurrency int) ports.DistributedWorkerpoolWorker {
	return &distributedWorkerpoolWorker{
		redisAddress: redisAddress,
		server: asynq.NewServer(
			asynq.RedisClientOpt{Addr: redisAddress},
			asynq.Config{
				Concurrency: concurrency,
			},
		),
		router: asynq.NewServeMux(),
	}
}

func (dww *distributedWorkerpoolWorker) RegisterTask(taskName string, handler func(ctx context.Context, t *ports.Task) error) {
	dww.router.HandleFunc(taskName, handler)
}

func (dww *distributedWorkerpoolWorker) ActiveWebConsole(monitoringAddress string) error {
	h := asynqmon.New(asynqmon.Options{
		RootPath:     "/",
		RedisConnOpt: asynq.RedisClientOpt{Addr: dww.redisAddress},
	})

	http.Handle(h.RootPath()+"/", h)

	return http.ListenAndServe(monitoringAddress, nil)
}

func (dww *distributedWorkerpoolWorker) StartWorker() error {
	err := dww.server.Run(dww.router)
	return err
}
