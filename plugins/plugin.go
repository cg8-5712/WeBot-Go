package plugins

import "github.com/eatmoreapple/openwechat"

type Plugin interface {
	Init() error
	Name() string
	HandleMessage(msg *openwechat.Message) error
}
