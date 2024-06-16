package bot

import (
	"os"

	traqwsbot "github.com/traPtitech/traq-ws-bot"
)

type BotService struct {
	bot           *traqwsbot.Bot
	postChannelID string // := "122c14e6-c32a-43b3-b905-d2aeb0c0a23e"
}

func NewBot() *BotService {
	bot, err := traqwsbot.NewBot(&traqwsbot.Options{
		AccessToken: os.Getenv("BOT_TOKEN"), // Required
	})
	if err != nil {
		panic(err)
	}
	return &BotService{bot: bot, postChannelID: "122c14e6-c32a-43b3-b905-d2aeb0c0a23e"}
}

func (s *BotService) Service() {
	if err := s.bot.Start(); err != nil {
		panic(err)
	}
}


func (s *BotService) PostNotify(title string, slideid string) {
	content := "New slide [**" + title + "**](https://h24s-04.trap.show/slides/" + slideid + ") has been posted :eyes:"
	s.BotSimplePost(s.postChannelID, content)
}
