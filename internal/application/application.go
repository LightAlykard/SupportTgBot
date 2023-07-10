package application

import (
	"context"

	"github.com/LightAlykard/SupportTgBot/internal/adapters/storage/pgstore"
	"github.com/LightAlykard/SupportTgBot/internal/bot"
	"github.com/LightAlykard/SupportTgBot/internal/config"
	"github.com/LightAlykard/SupportTgBot/internal/log"
	"github.com/LightAlykard/SupportTgBot/internal/repos/info"
)

func Start(ctx context.Context) {
	cfg, err := config.ReadConfig(".", ".env", "env")
	if err != nil {
		log.Error().Err(err).Send()
		panic(err)
	}

	var ist info.InfoStore

	dsn := cfg.DefaultBD
	pgst, err := pgstore.NewInfos(dsn)
	if err != nil {
		log.Error().Err(err).Send()
		panic(err)
	}
	defer pgst.Close()
	ist = pgst

	err = bot.InitBot(cfg, &ist)
	if err != nil {
		log.Error().Err(err).Send()
		panic(err)
	}
}

func Stop(ctx context.Context) {

}
