package services

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	TestServiceProvider,

// SERVICE_PROVIDER 锚点请勿删除! Do not delete this line!
)
