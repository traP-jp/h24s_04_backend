package bot

import (
	"os"

	traqwsbot "github.com/traPtitech/traq-ws-bot"
)

type BotService struct {
	bot *traqwsbot.Bot
}

func NewBot() *BotService {
	bot, err := traqwsbot.NewBot(&traqwsbot.Options{
		AccessToken: os.Getenv("BOT_TOKEN"), // Required
	})
	if err != nil {
		panic(err)
	}
	return &BotService{bot: bot}
}

func (s *BotService) Service() {

	if err := s.bot.Start(); err != nil {
		panic(err)
	}

}

// ここにPOSTした時の挙動を書く 引数にtitleとか入れておくと良い

func (s *BotService) PostNotify() {

}
