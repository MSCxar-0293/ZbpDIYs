//Package jrrp 简单的测人品
package jrrp

import (
	"math/rand"
	"time"
	"strconv"

	control "github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"

	"github.com/FloatTech/zbputils/control/order"
)

func init(){
	engine := control.Register("jrrp",order.AcquirePrio(), &control.Options{
		DisableOnDefault: false,
		Help: "jrrp\n" +
		"- 今日人品",
	})

			now := time.Now().Format("20060102")
			var signTF map[string](int)
			signTF = make(map[string](int))
			var result map[int64](int)
			result =  make(map[int64](int))
	engine.OnFullMatch ("今日人品").SetBlock(true).
	Handle(func(ctx *zero.Ctx){
			user := ctx.Event.UserID
			userS := strconv.FormatInt(user,10)
			var si string = now + userS
			rand.Seed(time.Now().UnixNano())
			today := rand.Intn(100)
		if signTF [si] == 0 {
			signTF [si] = (1)
			result [user] = (today)
		ctx.SendChain(message.At(user),message.Text(" 阁下今日的人品值为",result [user],"呢~\n"),message.Image("https://img.qwq.nz/images/2022/04/04/aab2985d94e996558b303be42a954a4f.jpg"))
		} else {
			ctx.SendChain(message.At(user),message.Text(" 阁下今日的人品值为",result [user],"呢~\n"),message.Image("https://img.qwq.nz/images/2022/04/04/aab2985d94e996558b303be42a954a4f.jpg"))
		}

	})
}
