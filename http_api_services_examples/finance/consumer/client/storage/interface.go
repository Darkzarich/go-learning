package storage

import "context"

type IStorageClient interface {
	SaveIntraday(ctx context.Context, intraday Intraday) error
}
