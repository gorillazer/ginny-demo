package repositories

import (
	"github.com/google/wire"

	mysql "github.com/gorillazer/ginny-mysql"

	redis "github.com/gorillazer/ginny-redis"
	// DATABASE_LIB 锚点请勿删除! Do not delete this line!
)

var ProviderSet = wire.NewSet(

	redis.Provider,
	mysql.Provider,
	// mongo.Provider,
	// DATABASE_PROVIDER 锚点请勿删除! Do not delete this line!
	UserRepositoryProvider,
	// REPO_PROVIDER 锚点请勿删除! Do not delete this line!
)
