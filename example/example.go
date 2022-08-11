// Package example
package example

import ( 
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init(){
	engine := control.Register("example", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Help: "example\n" +
		"- hello, world! \n" +
		"- 使用test以验证你的第一个插件。",
	})
	engine.OnFullMatch ("test").SetBlock(true).
	Handle(func(ctx *zero.Ctx){
		ctx.SendChain(message.Text("hello, world! "))
	})
	engine.OnFullMatch ("hello, world!").SetBlock(true).
	Handle(func(ctx *zero.Ctx){
		ctx.SendChain(message.Text("成功！"))
	})
}
