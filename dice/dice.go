//Package dice 简单骰子
package dice

import (
	"strconv"
	"math/rand"

	control "github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init(){
	engine := control.Register("dice", &control.Options{
		DisableOnDefault: false,
		Help: "dice\n" +
		"- 。r[骰子数量]d[骰子面数]\n" +
		"注意：无默认值！请不要省略任一参数。",
	})
	engine.OnRegex ("。r(.*)d(.*)").SetBlock(true).
	Handle(func(ctx *zero.Ctx){
		r1 := ctx.State["regex_matched"].([]string)[1]
		d1 := ctx.State["regex_matched"].([]string)[2]
		r, _ := strconv.Atoi(r1)
		d, _ := strconv.Atoi(d1)
		if r <= 100 && d <= 100 {
		res := rd (r,d)
		ctx.SendChain(message.At(ctx.Event.UserID), message.Text(" 阁下掷出了", res, "呢~"))
		} else {
		ctx.SendChain(message.Text("骰子太多啦~~数不过来了！"))
		}
	})
}

func rd(r, d int) string {
	sum := 0
	time := 0
	text := ""
	for time < r {
		time = time + 1
		res := rand.Intn(d)
		for res == 0 {
			res = rand.Intn(d)
		}
		sum += res
		resT := strconv.Itoa(res)
		sumT := strconv.Itoa(sum)
		var pre string
		if time == 1 {
			pre = ""
		} else {
			pre = "+"
		}
		text = text + pre + resT
		if time == r {
			text = text + "=" + sumT
		}
	}
	return text
}
