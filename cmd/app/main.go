package main

import (
	"context"
	"encoding/json"
	"internal/config"
	"internal/infrastructure"
	"internal/model"
	"internal/web"
	"log"

	"github.com/sethvargo/go-envconfig"
)

func main() {
	ctx := context.Background()

	var cfg config.AppConfig

	if err := envconfig.Process(ctx, &cfg); err != nil {
		log.Default().Fatal(err)
	}

	redis := infrastructure.NewRedis(cfg.RedisHostName, cfg.RedisPassword, 0)
	infrastructure.StartConsume(*redis, cfg.ConsumeChannel, ctx, func(s string) {
		var data model.SuricataMessageData
		if err := json.Unmarshal([]byte(s), &data); err != nil {
			return
		}
		web.CreateTheHiveAlert(data, cfg.ApiAddr, cfg.ApiToken)
	})
}
