package storage

import (
	"context"
	di "consumer/internal/domain/intraday"
)

type StorageClient interface {
	SaveIntraday(ctx context.Context, intraday di.Intraday) error
}
