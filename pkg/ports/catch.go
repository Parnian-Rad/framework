package ports

import (
	"context"
	"time"
)

type Catch interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, data interface{}, exp time.Duration) error
}
