// Package bilibiliforward Bilibili站内私信转发服务
package bilibiliforward

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"time"

	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

//一天 = 86400000000毫秒

func init() {
	engine := control.Register("bilibiliforward", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Help: "bilibiliforward\n" +
			"- 绑定cookie：+(cookie)\n" +
			"- 检查私信列表（注：默认过去一天内）\n" +
			"- 查看+(uid)+#+(数量)条消息\n" +
			"- 对+(uid)+说+(内容)\n" +
			"【注：cookie有效期不稳定，操作失败请尝试替换cookie】",
	})
	engine.OnFullMatch("检查私信列表", zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			var allteg AllTogether
			var msgcon MessageCon
			msglst := ""
			timesw := BeginOfTheDay() - 86400000000
			cookieInDB := GetCookieOf(ctx.Event.UserID)
			bodyText := ReceiveList(cookieInDB, timesw)
			err := json.Unmarshal(bodyText, &allteg)
			if err != nil {
				panic(err)
			}
			if allteg.Code == 0 {
				for _, session := range allteg.Data.Sessions {
					msgtime := dateTrans(session.LastMsg.TimeStamp) //获取最后一条消息发送时的时间
					msgrow := "送信人：" + GetUserName(cookieInDB, str64(session.TalkerID)) + "(" + str64(session.TalkerID) + ")\n最后发送时间：" + msgtime + "\n未读条数：" + str(session.UnreadCount) + "\n"
					err = json.Unmarshal([]byte(session.LastMsg.Content), &msgcon)
					if err != nil {
						fmt.Println(err)
					}
					if session.GrpName != "" || GetUserName(cookieInDB, str64(session.TalkerID)) == "" || str64(session.LastMsg.SenderUID) == ExtractUID(GetCookieOf(ctx.Event.UserID)) || session.UnreadCount == 0 { //跳过“应援团”（群聊）消息、跳过系统消息、跳过自己发送的消息导致的产生窗口、跳过零未读窗口
						continue
					}
					switch session.LastMsg.MsgType { //判断消息类型，1为普通私信，11为关注的up主视频发布推送 ...
					case 1: //普通文字消息
						msgrow = msgrow + "内容：" + msgcon.Content + "\n"
					case 2: //图片消息
						msgrow = msgrow + "内容：[图片消息]\n"
					case 6: //动图消息
						msgrow = msgrow + "内容：[动画表情]\n"
					case 7: //视频分享类消息
						msgrow = msgrow + "内容：[视频分享；" + msgcon.Title + "]\n"
					case 10: //获赠装扮类通知
						msgrow = msgrow + "内容：[" + msgcon.Title + "]\n"
					case 11: //关注的up发送的视频推送
						msgrow = msgrow + "~该消息为视频推送~\n视频标题：" + msgcon.Title + "\nBV：https://www.bilibili.com/video/" + msgcon.BVID + "\n" + str64(msgcon.Views) + " 播放 " + str64(msgcon.Danmks) + " 弹幕\n"
					}
					fmt.Println("Max Seqno:", session.MaxSeqno)
					msglst = msglst + msgrow + "-----------------------\n"
				}
				ctx.SendChain(message.Text("阁下过去一天内新增的消息如下：\n-----------------------\n" + msglst + "以上！"))
			} else {
				ctx.SendChain(message.Text("阁下...私信列表获取失败了呢...\n提示：" + allteg.Message))
			}
		})

	engine.OnRegex(`绑定cookie：(.*)`).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			usercookie := ctx.State["regex_matched"].([]string)[1] //获取信息中cookie
			uid := ctx.Event.UserID
			err := InsertCookieOf(uid, usercookie)
			if err != nil {
				panic(err)
			}
			cookieInDB := GetCookieOf(uid)
			ctx.SendChain(message.Text("成功！阁下当前的cookie为：" + cookieInDB))
		})
	engine.OnRegex(`查看(.*)#(.*)条消息`).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			seekuid := ctx.State["regex_matched"].([]string)[1]   //获取被读取的对话来源
			msgamount := ctx.State["regex_matched"].([]string)[2] //获取用户想要查看几条消息
			var sign string
			var uidall UIDAll
			var msgcon MessageConSingle
			uid := ctx.Event.UserID
			gid := ctx.Event.GroupID
			botid := ctx.Event.SelfID
			botname := zero.BotConfig.NickName[0]
			uidcookie := GetCookieOf(uid)
			usernn := GetUserName(uidcookie, seekuid)   //消息发送者b站昵称
			userpic := GetUserPhoto(uidcookie, seekuid) //消息发送者b站头像图链
			bodyText := GetUIDMsgs(GetCookieOf(uid), seekuid, msgamount)
			err := json.Unmarshal(bodyText, &uidall)
			if err != nil {
				panic(err)
			}
			if uidall.Code == 0 {
				msg := make(message.Message, str2int(msgamount))
				msg = append(msg, message.CustomNode(botname, botid, message.Message{message.Image(userpic), message.Text("\n~以下为与 " + usernn + " 的消息~")}))
				for _, msgdtd := range uidall.Data.Messages {
					err = json.Unmarshal([]byte(msgdtd.Content), &msgcon)
					if err != nil {
						fmt.Println(err)
					}
					if str64(msgdtd.SenderUID) == ExtractUID(GetCookieOf(uid)) { //判断是否为我方发送的消息
						sign = "【我方】"
					} else {
						sign = "【对方】"
					}
					usernn = GetUserName(uidcookie, str64(msgdtd.SenderUID)) //消息发送者b站昵称
					switch msgdtd.MsgType {
					case 1: //普通文字消息
						msg = append(msg, message.CustomNode(usernn, botid, message.Message{message.Text(sign + "[" + dateTrans(msgdtd.TimeStamp) + "]\n" + msgcon.Content)}))
					case 2: //图片类
						msg = append(msg, message.CustomNode(usernn, botid, message.Message{message.Text(sign + "[" + dateTrans(msgdtd.TimeStamp) + "]\n"), message.Image(msgcon.URL)}))
					case 6: //动图类
						msg = append(msg, message.CustomNode(usernn, botid, message.Message{message.Text(sign + "[" + dateTrans(msgdtd.TimeStamp) + "]\n"), message.Image(msgcon.URL)}))
					case 7: //视频分享类
						if msgcon.Title == "" {
							msg = append(msg, message.CustomNode(usernn, botid, message.Message{message.Text(sign + "[" + dateTrans(msgdtd.TimeStamp) + "][视频分享]\n*视频已删除*")}))
						} else {
							msg = append(msg, message.CustomNode(usernn, botid, message.Message{message.Text(sign + "[" + dateTrans(msgdtd.TimeStamp) + "][视频分享]\n"), message.Image(msgcon.SCoverLink), message.Text("\n标题：" + msgcon.Title + "\nBV：https://www.bilibili.com/video/" + msgcon.BVID)}))
						}
					case 10: //装扮赠送通知类
						var jumptext string
						if msgcon.JumpText != "" {
							jumptext = "\n*请前往客户端内点击“" + msgcon.JumpText + "”哦~"
						} else {
							jumptext = ""
						}
						msg = append(msg, message.CustomNode(usernn, botid, message.Message{message.Text(sign + "[" + dateTrans(msgdtd.TimeStamp) + "][" + msgcon.Title + "]\n" + msgcon.Text + jumptext)}))
					case 11: //视频推送
						if msgcon.Title == "" {
							msg = append(msg, message.CustomNode(usernn, botid, message.Message{message.Text(sign + "[" + dateTrans(msgdtd.TimeStamp) + "][视频推送]\n*视频已删除*")}))
						} else {
							msg = append(msg, message.CustomNode(usernn, botid, message.Message{message.Text(sign + "[" + dateTrans(msgdtd.TimeStamp) + "][视频推送]\n"), message.Image(msgcon.CoverLink), message.Text("\n标题：" + msgcon.Title + "\nBV：https://www.bilibili.com/video/" + msgcon.BVID + "\n" + str64(msgcon.Views) + " 播放 " + str64(msgcon.Danmks) + " 弹幕")}))
						}
					}
				}
				if gid != 0 {
					ctx.SendGroupForwardMessage(gid, msg)
				} else {
					ctx.SendPrivateForwardMessage(uid, msg)
				}
			} else {
				ctx.SendChain(message.Text("阁下...私信内容获取失败了呢...\n提示：" + uidall.Message))
			}
		})
	engine.OnRegex(`对(.*)说(.*)`).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			var picsum int
			send2uid := ctx.State["regex_matched"].([]string)[1] //获取信息中发送对象
			msginall := ctx.State["regex_matched"].([]string)[2] //获取信息中发送内容
			uid := ctx.Event.UserID
			cookieInDB := GetCookieOf(uid)
			imagelinks := extractImageLink(msginall)
			if len(imagelinks) != 0 {
				for i, link := range imagelinks {
					result := SendPic2U(cookieInDB, send2uid, link)
					if result == "0" {
						picsum = picsum + i
					} else {
						ctx.SendChain(message.Text("图片\n"), message.Image(link), message.Text("\n发送失败！\n错误信息："+result))
					}
				}
				ctx.SendChain(message.At(uid), message.Text(" 阁下成功发送了 "+str(picsum)+" 张图片√"))
			}
			sampleRegexp := regexp.MustCompile(`\[CQ:image(.*?),url=(.*?)\]`)
			purified := sampleRegexp.ReplaceAllString(msginall, "")
			result := SendText2U(cookieInDB, send2uid, purified)
			if result == "0" {
				ctx.SendChain(message.At(uid), message.Text(" 阁下的信息已经成功发送√"))
			} else {
				ctx.SendChain(message.At(uid), message.Text("阁下的信息发送失败了！\n错误信息："+result))
			}
		})
}

func dateTrans(timestamp int64) string {
	var thezero string
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	year := str(timeObj.Year())        //年
	month := str(int(timeObj.Month())) //月
	day := str(timeObj.Day())          //日
	hour := str(timeObj.Hour())        //小时
	minute := str(timeObj.Minute())    //分钟
	if len(minute) < 2 {
		thezero = "0"
	} else {
		thezero = ""
	}
	result := year + "/" + month + "/" + day + " " + hour + ":" + thezero + minute
	return result
}

func str64(num int64) string {
	a := strconv.FormatInt(num, 10)
	return a
}

func str(num int) string {
	a := strconv.Itoa(num)
	return a
}

func str2int(str string) int {
	a, _ := strconv.Atoi(str)
	return a
}

func BeginOfTheDay() int64 {
	timeStr := time.Now().Format("2006-01-02")
	loc, _ := time.LoadLocation("Asia/Hong_Kong")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, loc)
	timeUnix := t.UnixMicro()
	return timeUnix
}
