// Package card21 21点
package card21

import (
	"sort"
	"strconv"

	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init() {
	engine := control.Register("card21", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Help: "card21\n" +
			"- 加入21点\n" +
			"- 开始21点\n" +
			"- 21点抽牌\n" +
			"- 21点停止抽牌\n" +
			"- 退出21点\n" +
			"注：最少2人即可开始游戏，但是人数越多越好玩捏~",
	})
	cardnum := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	cardflower := []string{"黑桃", "红桃", "方块", "梅花"}
	players := []string{}
	carddeck := []string{}    // 总牌堆
	privatedeck := []string{} // 个人牌堆
	isplaying := 0            // 正在进行游戏判断器，重启后当前游戏作废
	plyord := 0
	scoreMap := make(map[string]int)

	engine.OnFullMatch("加入21点").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			if isplaying == 0 { // 判断当前是否有游戏，没有则可加入游戏
				players = append(players, strconv.FormatInt(ctx.Event.UserID, 10))
			} else {
				ctx.SendChain(message.Text("现在已经有一场21点游戏正在进行啦，阁下请好好观战哦~"))
			}
			ctx.SendChain(message.At(ctx.Event.UserID), message.Text(" 阁下成功加入游戏！\n当前人数 ", len(players), " 人。\n最少2人即可【开始21点】了哟~"))
		})
	engine.OnFullMatch("开始21点").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			if len(players) >= 2 { //判断人数是否达标
				isplaying = 1 // 标记游戏开始
				playerlist := ""
				for i := 0; i <= len(players)-1; i = i + 1 {
					playerlist = playerlist + ctx.CardOrNickName(puid(players[i])) + "(" + players[i] + ")\n"
				}
				ctx.SendChain(message.Text("さあ、ゲーム始める！\n玩家列表：\n"+playerlist+"\n请各位玩家依次发送【21点抽牌】进行抽牌！\n"), message.At(puid(players[0])), message.Text(" 阁下，从你开始~⭐"))
			} else {
				ctx.SendChain(message.Text("游戏人数还不够呢...\n当前人数 ", len(players), " 人。\n最少2人即可开始游戏了哟~"))
			}
		})
	engine.OnFullMatch("21点抽牌").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			card := cardflower[rdnum(4)] + cardnum[rdnum(13)] // 生成一张卡牌
			for Contains(carddeck, card) != -1 {              //判断是否生成重复卡牌，重复则重新生成
				card = cardflower[rdnum(4)] + cardnum[rdnum(13)]
			}
			if isplaying == 1 && ctx.Event.UserID == puid(players[plyord]) { // 判断是否正在游戏且是正确的玩家发言
				carddeck = append(carddeck, card)       // 总牌堆记牌
				privatedeck = append(privatedeck, card) // 私人牌堆记牌
				ctx.SendChain(message.At(ctx.Event.UserID), message.Text("阁下抽到了 "+card+" ！\n是否继续抽牌？\n> 【21点抽牌】\n> 【21点停止抽牌】"))
			} else if isplaying == 0 {
				ctx.SendChain(message.Text("当前并没有游戏呢~请阁下发送【加入21点】加入游戏吧~"))
			} else {
				ctx.SendChain(message.Text("现在还没有轮到阁下抽牌哦~请安静观战~~"))
			}
		})
	engine.OnFullMatch("21点停止抽牌").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			if or(judger17A(privatedeck) >= 0, judger17B(privatedeck) >= 0) && ctx.Event.UserID == puid(players[plyord]) {
				cardlist := ""
				for n := 0; n <= len(privatedeck)-1; n = n + 1 {
					cardlist = cardlist + privatedeck[n] + " "
				}
				if judger21A(privatedeck) == 0 || judger21B(privatedeck) == 0 {
					scoreMap[players[plyord]] = 100
					ctx.SendChain(message.Text("抽卡结束！\n 阁下手中的牌有：", cardlist, "\n点数达到21点，达成胜利！\n恭喜 "), message.At(ctx.Event.UserID), message.Text(" 阁下！"))
				} else if judger21A(privatedeck) > 0 && judger21B(privatedeck) > 0 {
					scoreMap[players[plyord]] = -100
					ctx.SendChain(message.Text("抽卡结束！\n 阁下手中的牌有：", cardlist, "\n点数超过21点，BOOM！\n残念， "), message.At(ctx.Event.UserID), message.Text(" 阁下输了！"))
				} else {
					scoreMap[players[plyord]] = judgerfin(judger21A(privatedeck), judger21B(privatedeck))
					ctx.SendChain(message.Text("抽卡结束！\n 阁下手中的牌有：", cardlist, "\n最优化点数为"+strconv.Itoa(judgerfin(judger21A(privatedeck), judger21B(privatedeck)))))
				}
				privatedeck = []string{}
				if plyord < len(players)-1 {
					plyord = plyord + 1
					ctx.SendChain(message.Text("有请下一位选手：\n"+ctx.CardOrNickName(puid(players[plyord]))+"("+players[plyord]+")\n"), message.At(puid(players[plyord])))
				} else if plyord == len(players)-1 {
					scorelist := ""
					for v := 0; v <= len(players)-1; v = v + 1 {
						scorelist = scorelist + ctx.CardOrNickName(puid(players[v])) + "(" + players[v] + ")：\n" + strconv.Itoa(scoreMap[players[v]]) + " 分\n"
					}
					ctx.SendChain(message.Text("游戏结束！成绩如下："+scorelist+"恭喜"), message.At(puid(winner(scoreMap))), message.Text("获得冠军！"))
					for m := 0; m <= len(players)-1; m = m + 1 {
						delete(scoreMap, players[m])
					}
					plyord = 0
					players = []string{}
					carddeck = []string{}
					isplaying = 0
				}
			} else if judger17A(privatedeck) < 0 && judger17B(privatedeck) < 0 && ctx.Event.UserID == puid(players[plyord]) {
				ctx.SendChain(message.Text("阁下当前的点数还未到达17点，还不能停止抽牌哦~\n请发送【21点抽牌】继续抽牌！"))
			} else {
				ctx.SendChain(message.Text("现在还没有轮到阁下抽牌哦~请安静观战~~"))
			}
		})
	engine.OnFullMatch("退出21点").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			sort.Strings(players)
			idindex := sort.SearchStrings(players, strconv.FormatInt(ctx.Event.UserID, 10))
			if idindex > len(players)-1 {
				ctx.SendChain(message.At(ctx.Event.UserID), message.Text(" 阁下还没有加入游戏呢......"))
			} else if isplaying == 1 {
				ctx.SendChain(message.Text("现在已经有一场21点游戏正在进行啦，阁下请好好观战哦~"))
			} else {
				players = append(players[:idindex], players[idindex+1:]...)
				ctx.SendChain(message.At(ctx.Event.UserID), message.Text(" 阁下成功退出游戏！\n当前人数 ", len(players), " 人。\n最少2人即可开始游戏了哟~"))
			}
		})
}
