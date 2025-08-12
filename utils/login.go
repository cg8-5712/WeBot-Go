package utils

import (
	"fmt"
	"log"

	"github.com/eatmoreapple/openwechat"
	"github.com/skip2/go-qrcode"
)

// ConsoleQrCode 终端打印二维码
func ConsoleQrCode(uuid string) {
	url := "https://login.weixin.qq.com/l/" + uuid
	qr, _ := qrcode.New(url, qrcode.Low)
	fmt.Println(qr.ToString(true))
}

// LoginWechatBot 根据配置登录，返回登录后的 Bot
func LoginWechatBot() *openwechat.Bot {
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式更稳定

	bot.UUIDCallback = ConsoleQrCode
	bot.ScanCallBack = func(resp openwechat.CheckLoginResponse) {
		avatar, _ := resp.Avatar()
		fmt.Println("扫码检测到头像:", avatar)
	}
	bot.LoginCallBack = func(_ openwechat.CheckLoginResponse) {
		fmt.Println("✅ 登录成功")
	}

	mode := Cfg.LoginMode
	storageFile := Cfg.HotReloadFile

	switch mode {
	case "hot":
		fmt.Println("🔄 使用热登录模式")
		reloadStorage := openwechat.NewFileHotReloadStorage(storageFile)
		defer reloadStorage.Close()

		if err := bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
			log.Println("热登录失败，回退扫码登录:", err)
			if err := bot.Login(); err != nil {
				log.Fatalf("扫码登录失败: %v", err)
			}
		}

	case "push":
		fmt.Println("📲 使用 PushLogin 模式")
		reloadStorage := openwechat.NewFileHotReloadStorage(storageFile)
		defer reloadStorage.Close()

		if err := bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
			log.Println("PushLogin 失败，回退扫码登录:", err)
			if err := bot.Login(); err != nil {
				log.Fatalf("扫码登录失败: %v", err)
			}
		}

	default:
		fmt.Println("📷 使用普通扫码登录模式")
		if err := bot.Login(); err != nil {
			log.Fatalf("扫码登录失败: %v", err)
		}
	}

	return bot
}
