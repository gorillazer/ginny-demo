package repo

import (
	"context"

	"github.com/google/wire"
	"github.com/goriller/ginny-demo/internal/cache"
	"github.com/goriller/ginny-demo/internal/config"
	"github.com/goriller/ginny-demo/internal/repo/entity"
	mysql "github.com/goriller/ginny-mysql"
	"github.com/goriller/ginny/logger"
	"go.uber.org/zap"
	// mongo "github.com/goriller/ginny-mongo"
	// DATABASE_LIB 锚点请勿删除! Do not delete this line!
)

// UserRepoProvider
var UserRepoProvider = wire.NewSet(NewUserRepo,
	wire.Bind(new(IUserRepo), new(*UserRepo)))

// IUserRepo
type IUserRepo interface {
	GetUser(ctx context.Context) (*entity.UserEntity, error)
}

// UserRepo
type UserRepo struct {
	config *config.Config
	mysql  *mysql.SqlBuilder
	// mongo  *mongo.Manager
	cache cache.IRedisCache
	// STRUCT_ATTR 锚点请勿删除! Do not delete this line!
}

// NewUserRepo
func NewUserRepo(
	config *config.Config,
	mysql *mysql.SqlBuilder,
	// mongo *mongo.Manager,
	cache cache.IRedisCache,
	// FUNC_PARAM 锚点请勿删除! Do not delete this line!
) *UserRepo {
	return &UserRepo{
		config: config,
		cache:  cache,
		mysql:  mysql,
		// mongo: mongo,
		// FUNC_ATTR 锚点请勿删除! Do not delete this line!
	}
}

func (p *UserRepo) GetUser(ctx context.Context) (*entity.UserEntity, error) {
	r := &entity.UserEntity{}
	log := logger.WithContext(ctx).With(zap.String("action", "GetUser"))
	// if err := p.mysql.Find(ctx, r, r.TableName(), nil); err != nil {
	// 	log.Error("", zap.Error(err))
	// 	return nil, err
	// }

	// if _, err := p.mongo.Database.Collection(r.TableName()).Find(ctx, nil); err != nil {
	// 	return nil, err
	// }
	// p.redis.DB().Get(ctx, r.TableName()).Result()
	log.Debug("GetUser")
	return r, nil
}
