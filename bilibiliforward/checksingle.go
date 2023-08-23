package bilibiliforward

import (
	"io"
	"log"
	"net/http"
)

type MessageConSingle struct {
	Content    string `json:"content"`
	Author     string `json:"author"`
	Title      string `json:"title"`
	CoverLink  string `json:"cover"`
	SCoverLink string `json:"thumb"`
	Describe   string `json:"desc"`
	BVID       string `json:"bvid"`
	ImageType  string `json:"imageType"`
	URL        string `json:"url"`
	Text       string `json:"text"`
	Views      int64  `json:"view"`
	Danmks     int64  `json:"danmaku"`
	PubDate    int64  `json:"pub_date"`
	JumpText   string `json:"jump_text"`
}

type MessageLST struct {
	SenderUID int64  `json:"sender_uid"`
	Content   string `json:"content"`
	MsgSeqno  int64  `json:"msg_seqno"`
	MsgType   int    `json:"msg_type"`
	TimeStamp int64  `json:"timestamp"`
}
type UIDData struct {
	Messages []MessageLST `json:"messages"` // 修改这里的标签
	HasMore  int          `json:"has_more"`
}

type UIDAll struct {
	Code    int     `json:"code"`
	Msg     string  `json:"msg"`
	Message string  `json:"message"`
	TTL     int     `json:"ttl"`
	Data    UIDData `json:"data"`
}

func GetUIDMsgs(cookie string, uid string, amount string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.vc.bilibili.com/svr_sync/v1/svr_sync/fetch_session_msgs?sender_device_id=1&talker_id="+uid+"&session_type=1&size="+amount+"&build=0&mobi_app=web", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "api.vc.bilibili.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("cookie", cookie)
	req.Header.Set("origin", "https://message.bilibili.com")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("referer", "https://message.bilibili.com/")
	req.Header.Set("sec-ch-ua", `"Chromium";v="116", "Not)A;Brand";v="24", "Google Chrome";v="116"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return bodyText
}
