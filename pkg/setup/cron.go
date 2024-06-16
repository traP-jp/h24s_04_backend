package setup

import (
	"h24s_04/pkg/transfer"
	"time"

	"github.com/robfig/cron/v3"
)

func Cronsetup(tr transfer.ITransferFileService) {

	c := cron.New()
	c.AddFunc("0-59/1 * * * *", tr.Urlupdate)

	c.Start()

	time.Sleep(time.Second * 5) //cronスタート用

}
