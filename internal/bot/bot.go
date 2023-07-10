package bot

import (
	"fmt"
	"net/http"

	"github.com/LightAlykard/SupportTgBot/internal/config"
	zlog "github.com/LightAlykard/SupportTgBot/internal/log"
	"github.com/LightAlykard/SupportTgBot/internal/repos/info"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

type bot struct {
	bot *tgbotapi.BotAPI
	db  *info.InfoStore
	cfg config.Config
}

func InitBot(cfg config.Config, db *info.InfoStore) error {
	botAPI, err := tgbotapi.NewBotAPI(cfg.TelegramLoggerBotToken)
	if err != nil {
		return err
	}

	botState := &bot{
		bot: botAPI,
		db:  db,
		cfg: cfg,
	}
	botState.InitUpdates()

	return nil
}

func (b *bot) InitUpdates() {

	b.bot.Debug = true

	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	wh, _ := tgbotapi.NewWebhookWithCert("https://www.example.com:8443/"+b.bot.Token, "cert.pem")

	_, err := b.bot.Request(wh)
	if err != nil {
		zlog.Error().Err(err).Send()
	}

	info, err := b.bot.GetWebhookInfo()
	if err != nil {
		zlog.Error().Err(err).Send()
	}

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
		//лог
	}

	updates := b.bot.ListenForWebhook("/" + b.bot.Token)
	go http.ListenAndServeTLS(fmt.Sprintf("%s:%s", b.cfg.BotAddress, b.cfg.BotPort), b.cfg.CertPath, b.cfg.KeyPath, nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}

	// messege := func(w http.ResponseWriter, r *http.Request) {
	// 	text, err := ioutil.ReadAll(r.Body)
	// 	//log
	// 	if err != nil {
	// 		zlog.Error().Err(err).Send()
	// 	}
	// 	var botText message.BotMessage
	// 	err = json.Unmarshal(text, &botText)
	// 	//log
	// 	fmt.Println(fmt.Sprintf("%s", text))

	// 	//распарс по переменным

	// 	//функц обработки ответа

	// }

	// http.HandleFunc("/", messege)
	// log.Fatal(http.ListenAndServeTLS(fmt.Sprintf("%s:%s", b.cfg.BotAddress, b.cfg.BotPort), b.cfg.CertPath, b.cfg.KeyPath, nil))

}
