package echo

import (
	"log"

	"WeBot/plugins"
	"github.com/eatmoreapple/openwechat"
)

type EchoPlugin struct{}

func (p *EchoPlugin) Init() error {
	log.Println("Echo 插件初始化完成")
	return nil
}

func (p *EchoPlugin) Name() string {
	return "Echo"
}

func (p *EchoPlugin) HandleMessage(msg *openwechat.Message) error {
	//if msg.IsText() {
	//	return msg.ReplyText("Echo: " + msg.Content)
	//}
	return nil
}

// 确保实现接口
var _ plugins.Plugin = (*EchoPlugin)(nil)
