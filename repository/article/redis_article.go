package article

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Fajar-Islami/ais_code_test/entity"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type redisArticleRepoImpl struct {
	context     context.Context
	redisClient *redis.Client
}

type RedisArtilceRepository interface {
	GetArticleByQueryCtx(key string) ([]entity.Article, error)
	SetArticleCtx(key string, time int, data *[]entity.Article) error
	DeleteArticleCtx(key string) error
}

func NewRedisRepo(context context.Context, redisClient *redis.Client) RedisArtilceRepository {
	return &redisArticleRepoImpl{
		context:     context,
		redisClient: redisClient,
	}
}

func (ra *redisArticleRepoImpl) GetArticleByQueryCtx(key string) ([]entity.Article, error) {
	fmt.Printf("Get keys %s from redis\n", key)
	newBytes, err := ra.redisClient.Get(ra.context, key).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "articleRedisRepo.GetArticleByQueryCtx.redisClient.Get")
	}

	newBase := &[]entity.Article{}
	if err = json.Unmarshal(newBytes, newBase); err != nil {
		return nil, errors.Wrap(err, "articleRedisRepo.GetArticleByQueryCtx.json.Unmarshal")
	}

	return *newBase, nil
}

func (ra *redisArticleRepoImpl) SetArticleCtx(key string, times int, data *[]entity.Article) error {
	newBytes, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "articleRedisRepo.SetArticleCtx.json.Marshal")
	}

	if err = ra.redisClient.Set(ra.context, key, newBytes, time.Minute*time.Duration(times)).Err(); err != nil {
		return errors.Wrap(err, "articleRedisRepo.SetArticleCtx.redisClient.set")
	}
	fmt.Printf("Set keys %s to redis\n", key)
	return nil
}

func (ra *redisArticleRepoImpl) DeleteArticleCtx(key string) error {
	if err := ra.redisClient.Del(ra.context, key).Err(); err != nil {
		return errors.Wrap(err, "articleRedisRepo.DeleteArticleCtx.redisClient.Del")
	}
	fmt.Printf("Delete keys %s from redis\n", key)
	return nil
}
