package intraday

import "errors"

var ErrRejected = errors.New("couldn't get intraday from storage")
