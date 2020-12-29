package repeat

import (
	"github.com/wdvxdr1123/ZeroBot"
)

func init() {
	zero.RegisterPlugin(testPlugin{}) // 注册插件
}

type testPlugin struct{}

func (_ testPlugin) GetPluginInfo() zero.PluginInfo { // 返回插件信息
	return zero.PluginInfo{
		Author:     "wdvxdr1123",
		PluginName: "test",
		Version:    "0.1.0",
		Details:    "这是一个测试复读插件",
	}
}

func (_ testPlugin) Start() { // 插件主体

	zero.OnCommand("echo").Handle(handleCommand)

	zero.OnCommand("cmd_echo").Handle(handleCommand)

	zero.OnSuffix("复读").Handle(func(_ *zero.Matcher, event zero.Event, state zero.State) zero.Response {
		zero.Send(event, state["args"])
		return zero.FinishResponse
	})

	zero.OnFullMatch("你是谁", zero.OnlyToMe).Handle(func(matcher *zero.Matcher, event zero.Event, state zero.State) zero.Response {
		zero.Send(event, "我是一个复读机~~~")
		echo := matcher.Get("我想要复读你的话!")
		zero.Send(event, echo)
		return zero.FinishResponse
	})
}

func handleCommand(_ *zero.Matcher, event zero.Event, state zero.State) zero.Response {
	zero.Send(event, state["args"])
	return zero.FinishResponse
}
