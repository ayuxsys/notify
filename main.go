package main

import (
	"github.com/ayuxsys/notify/config"
	"github.com/ayuxsys/notify/log"
	"github.com/ayuxsys/notify/notifier"
)

func init() {
	config.Setup()
}

func main() {
	c := &cmdi{}
	c.parseFlags()
	c.Run()
}

func (c *cmdi) Run() {
	notify := notifier.New(config.Load(c.cfgFile))
	notify.File = c.notify.file
	notify.Msg = c.notify.msg

	if err := notify.Telegram(); err != nil {
		log.Error.Fatal(err)
	}
}
