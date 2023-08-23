package bilibiliforward

import (
	"io"
	"log"
	"net/http"
)

type MessageCon struct {
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
}

type LastMsg struct {
	SenderUID int64  `json:"sender_uid"`
	MsgType   int    `json:"msg_type"`
	Content   string `json:"content"`
	TimeStamp int64  `json:"timestamp"`
}

type SessionList struct {
	TalkerID    int64   `json:"talker_id"`
	UnreadCount int     `json:"unread_count"`
	LastMsg     LastMsg `json:"last_msg"`
	GrpName     string  `json:"group_name"`
	MaxSeqno    int64   `json:"max_seqno"`
}
type Data struct {
	Sessions []SessionList `json:"session_list"`
	HasMore  int           `json:"has_more"`
}

type AllTogether struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    Data   `json:"data"`
}

func ReceiveList(cookie string, begintime int64) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.vc.bilibili.com/session_svr/v1/session_svr/new_sessions?begin_ts="+str64(begintime)+"&build=0&mobi_app=web", nil)
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
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="99", "Google Chrome";v="115", "Chromium";v="115"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
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
