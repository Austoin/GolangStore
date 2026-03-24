package redis

import (
	"context"
	"fmt"

	"github.com/austoin/GolangStore/pkg/config"
	githubredis "github.com/redis/go-redis/v9"
)

func BuildAddr(conf config.Redis) string {
	return fmt.Sprintf("%s:%s", conf.Host, conf.Port)
}

func NewClient(conf config.Redis) *githubredis.Client {
	return githubredis.NewClient(&githubredis.Options{
		Addr:     BuildAddr(conf),
		Password: conf.Password,
	})
}

func Ping(ctx context.Context, client *githubredis.Client) error {
	return client.Ping(ctx).Err()
}
