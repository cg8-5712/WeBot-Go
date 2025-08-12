package main

import (
	"os"
	"os/signal"
	"syscall"

	"WeBot/plugins/echo"
	"WeBot/utils"
	"github.com/eatmoreapple/openwechat"
)

func main() {
	utils.LoadConfig()
	utils.InitLogLevelFromConfig()

	utils.Info("日志模块初始化完成，当前日志级别: %s", utils.Cfg.LogLevel)

	bot := utils.LoginWechatBot()
	if bot == nil {
		utils.Error("登录微信机器人失败")
	}
	utils.Info("微信机器人登录成功")

	// 初始化插件（后续可动态加载多个插件）
	echoPlugin := &echo.EchoPlugin{}
	if err := echoPlugin.Init(); err != nil {
		utils.Error("插件初始化失败: %v", err)
	}
	utils.Info("Echo 插件初始化成功")

	// 注册消息处理器
	bot.MessageHandler = func(msg *openwechat.Message) {
		utils.Debug("接收到消息: %v", msg.Content)
		if err := echoPlugin.HandleMessage(msg); err != nil {
			utils.Error("处理消息出错: %v", err)
		}
	}

	// 优雅退出处理
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		utils.Info("收到退出信号，开始注销微信账号...")
		if err := bot.Logout(); err != nil {
			utils.Error("注销失败: %v", err)
		} else {
			utils.Info("注销成功")
		}
		os.Exit(0)
	}()

	utils.Info("微信机器人启动成功，等待消息...")
	bot.Block()
}
