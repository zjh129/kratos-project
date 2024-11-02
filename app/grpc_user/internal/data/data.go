package data

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/mgtv-tech/jetcache-go/local"
	"github.com/mgtv-tech/jetcache-go/remote"
	"kratos-project/app/grpc_user/internal/conf"
	"kratos-project/app/grpc_user/internal/data/ent"
	"kratos-project/app/grpc_user/internal/data/ent/migrate"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	cache "github.com/mgtv-tech/jetcache-go"
	"github.com/redis/go-redis/v9"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
	_ "kratos-project/app/grpc_user/internal/data/ent/runtime"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClient, NewJetCache)

var ErrRecordNotFound = errors.New("record not found")

// Data .
type Data struct {
	db    *ent.Client
	cache cache.Cache
	log   *log.Helper
}

// NewData .
func NewData(
	c *conf.Data,
	logger log.Logger,
	entClient *ent.Client,
	jetCache cache.Cache,
) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "kratos-learn/data"))
	d := &Data{
		db:    entClient,
		cache: jetCache,
		log:   log,
	}
	cleanup := func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}
	return d, cleanup, nil
}

func NewEntClient(c *conf.Data, logger log.Logger) *ent.Client {
	log := log.NewHelper(log.With(logger, "module", "kratos-learn/data/ent"))

	client, err := ent.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

// NewJetCache init jetcache
func NewJetCache(c *conf.Data, logger log.Logger) cache.Cache {
	log := log.NewHelper(log.With(logger, "module", "kratos-learn/data/jetcache"))
	redisAddrs := strings.Split(c.Redis.Addr, ":")
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			redisAddrs[0]: ":" + redisAddrs[1],
		},
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
	})
	sourceID := uuid.NewString() // Unique identifier for this cache instance
	channelName := "syncUserChannel"
	pubSub := ring.Subscribe(context.Background(), channelName)
	mycache := cache.New(cache.WithName("kratos-learn"),
		cache.WithRemote(remote.NewGoRedisV9Adapter(ring)),
		cache.WithLocal(local.NewFreeCache(256*local.MB, time.Minute)),
		cache.WithErrNotFound(ErrRecordNotFound),
		cache.WithRefreshDuration(time.Minute),
		cache.WithSyncLocal(true),
		cache.WithEventHandler(func(event *cache.Event) {
			// Broadcast local cache invalidation for the received keys
			bs, _ := json.Marshal(event)
			ring.Publish(context.Background(), channelName, string(bs))
		}),
	)
	go func() {
		for {
			msg := <-pubSub.Channel()
			var event *cache.Event
			if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
				panic(err)
			}
			log.Infof("Received event: %v", event)

			// Invalidate local cache for received keys (except own events)
			if event.SourceID != sourceID {
				for _, key := range event.Keys {
					mycache.DeleteFromLocalCache(key)
				}
			}
		}
	}()

	return mycache
}
