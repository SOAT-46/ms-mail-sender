package listeners

import "github.com/google/wire"

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	NewSendSuccessMailQueueListener,
	NewSendErrorMailQueueListener,
)
