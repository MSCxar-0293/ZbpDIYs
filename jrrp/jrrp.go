//Package jrrp 简单的测人品
package jrrp

import (
	"math/rand"
	"time"
	"strconv"

	control "github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init(){
	engine := control.Register("jrrp", &control.Options{
		DisableOnDefault: false,
		Help: "jrrp\n" +
		"- 今日人品",
	})
			var signTF map[string](int)
			signTF = make(map[string](int))
			var result map[string](int)
			result =  make(map[string](int))
	engine.OnFullMatch ("今日人品").SetBlock(true).
	Handle(func(ctx *zero.Ctx){
		        now := time.Now().Format("20060102")
			user := ctx.Event.UserID
			userS := strconv.FormatInt(user,10)
			var si string = now + userS
		if signTF [si] == 0 {
			signTF [si] = (1)
			today := rand.Intn(100)
			result [si] = (today)
		ctx.SendChain(message.At(user),message.Text(" 阁下今日的人品值为",result [si],"呢~\n"),message.Image("https://img.qwq.nz/images/2022/04/04/aab2985d94e996558b303be42a954a4f.jpg"))
		} else {
			ctx.SendChain(message.At(user),message.Text(" 阁下今日的人品值为",result [si],"呢~\n"),message.Image("https://img.qwq.nz/images/2022/04/04/aab2985d94e996558b303be42a954a4f.jpg"))
		}

	})
}
