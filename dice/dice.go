//Package dice 简单骰子
package dice

import (
	"strconv"
	"math/rand"

	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init(){
	engine := control.Register("dice", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Help: "dice\n" +
		"- 。r[骰子数量]d[骰子面数]\n" +
		"- 。ra[属性名称]#[次数]#[成功率]\n" ,
	})
	engine.OnRegex ("。r(.*)d(.*)", zero.OnlyToMe).SetBlock(true).
	Handle(func(ctx *zero.Ctx){
		r1 := ctx.State["regex_matched"].([]string)[1]
		d1 := ctx.State["regex_matched"].([]string)[2]
		if r1 == "" {
			r1 = "1"
		}
		if d1 == "" {
			d1 = "100"
		}
		r, _ := strconv.Atoi(r1)
		d, _ := strconv.Atoi(d1)
		if d == 1 || d == 0 || r == 0 {
		ctx.SendChain(message.Text("阁下..你在让我骰什么啊？( ´_ゝ`)"))
		} else {
		if r <= 100 && d <= 100 {
		res := rd (r,d)
		ctx.SendChain(message.At(ctx.Event.UserID), message.Text(" 阁下掷出了", res, "呢~"))
		} else {
		ctx.SendChain(message.Text("骰子太多啦~~数不过来了！"))
		}
		}
	})
	engine.OnRegex ("。ra(.*)#(.*)#(.*)").SetBlock(true).
	Handle(func(ctx *zero.Ctx){
		text := ctx.State["regex_matched"].([]string)[1]
		timeT := ctx.State["regex_matched"].([]string)[2]
		a1 := ctx.State["regex_matched"].([]string)[3]
		rate, _ := strconv.Atoi(a1)
		times, _ := strconv.Atoi(timeT)
		ima := 0
		msg := ""
		for ima < times {
			ima = ima + 1
			res := ra(rate)
			next := "\n"
			if ima == times {
				next = ""
			}
			msg = msg + res + next
		}
		ctx.SendChain(message.At(ctx.Event.UserID), message.Text("阁下进行", text, "的结果是：\n", msg))
		})
}

func rd(r, d int) string {
	sum := 0
	time := 0
	rT := strconv.Itoa(r)
	dT := strconv.Itoa(d)
	text := "R" + rT + "D" + dT + "="
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

func ra(rate int) string {
	res := rand.Intn(100)
	for res == 0 {
		res = rand.Intn(100)
	}
	resT := strconv.Itoa(res)
	rateT := strconv.Itoa(rate)
	text := "D100=" + resT + "/" + rateT + " ~"
	if res == 100 {
		text = text + "大☆失☆败"
		return text
	}
	if res == 1 {
		text = text + "大☆成☆功"
		return text
	}
	if res <= rate / 5 {
		text = text + "极难成功"
		return text
	}
	if res <= rate / 2 {
		text = text + "困难成功"
		return text
	}
	if res <= rate {
		text = text + "成功"
		return text
	}
	if rate >= 50 || res < 96 {
		text = text + "失败"
		return text
	}
	text = text + "大☆失☆败"
	return text
}
