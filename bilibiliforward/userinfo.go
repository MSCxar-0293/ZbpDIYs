package bilibiliforward

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Card struct {
	UID     string `json:"mid"`
	Name    string `json:"name"`
	ProfPic string `json:"face"`
}

type InfoBasic struct {
	Code int  `json:"code"`
	Card Card `json:"card"`
}

func GetUserName(cookie string, uid string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://account.bilibili.com/api/member/getCardByMid?mid="+uid, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "account.bilibili.com")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("cookie", cookie)
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("sec-ch-ua", `"Chromium";v="116", "Not)A;Brand";v="24", "Google Chrome";v="116"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
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
	var infobsc InfoBasic
	err = json.Unmarshal(bodyText, &infobsc)
	if err != nil {
		panic(err)
	}
	username := infobsc.Card.Name
	return username
}

func GetUserPhoto(cookie string, uid string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://account.bilibili.com/api/member/getCardByMid?mid="+uid, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "account.bilibili.com")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("cookie", cookie)
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("sec-ch-ua", `"Chromium";v="116", "Not)A;Brand";v="24", "Google Chrome";v="116"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
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
	var infobsc InfoBasic
	err = json.Unmarshal(bodyText, &infobsc)
	if err != nil {
		panic(err)
	}
	userphoto := infobsc.Card.ProfPic
	return userphoto
}
