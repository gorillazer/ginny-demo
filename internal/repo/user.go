package repo

import (
	"context"

	"github.com/google/wire"
	"github.com/goriller/ginny-demo/internal/repo/entity"
	mysql "github.com/goriller/ginny-mysql"
	// mongo "github.com/goriller/ginny-mongo"
	// redis "github.com/goriller/ginny-redis"
	// DATABASE_LIB 锚点请勿删除! Do not delete this line!
)

// UserRepositoryProvider
var UserRepositoryProvider = wire.NewSet(NewUserRepository,
	wire.Bind(new(IUserRepository), new(*UserRepository)))

// IUserRepository
type IUserRepository interface {
	GetUser(ctx context.Context) (*UserRepository, error)
}

// UserRepository
type UserRepository struct {
	// redis *redis.Manager
	mysql *mysql.SqlBuilder
	// mongo *mongo.Manager
	// STRUCT_ATTR 锚点请勿删除! Do not delete this line!
}

// NewUserRepository
func NewUserRepository(
	// redis *redis.Manager,
	mysql *mysql.SqlBuilder,
	// mongo *mongo.Manager,
	// FUNC_PARAM 锚点请勿删除! Do not delete this line!
) *UserRepository {
	return &UserRepository{

		// redis: redis,
		mysql: mysql,
		// mongo: mongo,
		// FUNC_ATTR 锚点请勿删除! Do not delete this line!
	}
}

func (p *UserRepository) GetUser(ctx context.Context) (*entity.UserEntity, error) {
	r := &entity.UserEntity{}
	// if err := p.mysql.Find(ctx, r, r.TableName(), nil); err != nil {
	// 	p.logger.Error("", zap.Error(err))
	// 	return nil, err
	// }

	// if _, err := p.mongo.Database.Collection(r.TableName()).Find(ctx, nil); err != nil {
	// 	return nil, err
	// }
	// p.redis.DB().Get(ctx, r.TableName()).Result()
	return r, nil
}
