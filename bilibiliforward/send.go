package bilibiliforward

import (
	"encoding/json"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type SendStat struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
}

func SendText2U(cookie string, uid string, content string) string {
	now := time.Now().Unix()
	msgcontent := "{\"content\":\"" + content + "\"}"
	encodedmsgc := url.QueryEscape(msgcontent)
	client := &http.Client{}
	var data = strings.NewReader("msg%5Bsender_uid%5D=" + ExtractUID(cookie) + "&msg%5Breceiver_id%5D=" + uid + "&msg%5Breceiver_type%5D=1&msg%5Bmsg_type%5D=1&msg%5Bmsg_status%5D=0&msg%5Bcontent%5D=" + encodedmsgc + "&msg%5Btimestamp%5D=" + str64(now) + "&msg%5Bnew_face_version%5D=0&msg%5Bdev_id%5D=A63F4006-DDBE-43C4-9285-3A842052735D&from_firework=0&build=0&mobi_app=web&csrf_token=" + ExtractJctToken(cookie) + "&csrf=" + ExtractJctToken(cookie))
	req, err := http.NewRequest("POST", "https://api.vc.bilibili.com/web_im/v1/web_im/send_msg", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "api.vc.bilibili.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
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
	var sendstat SendStat
	err = json.Unmarshal(bodyText, &sendstat)
	if err != nil {
		panic(err)
	}
	return sendstat.Message
}

func SendPic2U(cookie string, uid string, piclink string) string {
	now := time.Now().Unix()
	x, y := SizeParse(piclink)
	xf := x / 2
	yf := y / 2
	msgcontent := "{\"url\":\"" + piclink + "\",\"height\":" + str(yf) + ",\"width\":" + str(xf) + ",\"imageType\":\"jpeg\"}"
	encodedmsgc := url.QueryEscape(msgcontent)
	client := &http.Client{}
	var data = strings.NewReader("msg%5Bsender_uid%5D=" + ExtractUID(cookie) + "&msg%5Breceiver_id%5D=" + uid + "&msg%5Breceiver_type%5D=1&msg%5Bmsg_type%5D=2&msg%5Bmsg_status%5D=0&msg%5Bcontent%5D=" + encodedmsgc + "&msg%5Btimestamp%5D=" + str64(now) + "&msg%5Bnew_face_version%5D=0&msg%5Bdev_id%5D=A63F4006-DDBE-43C4-9285-3A842052735D&from_firework=0&build=0&mobi_app=web&csrf_token=" + ExtractJctToken(cookie) + "&csrf=" + ExtractJctToken(cookie))
	req, err := http.NewRequest("POST", "https://api.vc.bilibili.com/web_im/v1/web_im/send_msg", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "api.vc.bilibili.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
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
	var sendstat SendStat
	err = json.Unmarshal(bodyText, &sendstat)
	if err != nil {
		panic(err)
	}
	return sendstat.Message
}

func ExtractUID(cookie string) string {
	//uid前后标识
	startIndex := strings.Index(cookie, "; DedeUserID=")
	endIndex := strings.Index(cookie, "; DedeUserID__ckMd5")

	if startIndex != -1 && endIndex != -1 {
		uidincookie := cookie[startIndex+len("; DedeUserID=") : endIndex]
		return uidincookie
	} else {
		a := "【未识别到cookie内uid】"
		return a
	}
}

func ExtractJctToken(cookie string) string {
	//该token前后标识
	startIndex := strings.Index(cookie, "; bili_jct=")
	endIndex := strings.Index(cookie, "; sid")

	if startIndex != -1 && endIndex != -1 {
		tokenincookie := cookie[startIndex+len("; bili_jct=") : endIndex]
		return tokenincookie
	} else {
		a := "【未识别到cookie内token】"
		return a
	}
}

func SizeParse(piclink string) (int, int) {
	res, err := http.Get(piclink)
	if err != nil || res.StatusCode != 200 {
		panic(err)
	}
	defer res.Body.Close()

	m, _, err := image.Decode(res.Body)
	if err != nil {
		panic(err)
	}

	x := m.Bounds().Dx() // 宽：32
	y := m.Bounds().Dy() // 高：32
	return x, y
}

func extractImageLink(text string) []string {
	re := regexp.MustCompile(`\[CQ:image(.*?),url=(.*?)\]`) //匹配用户发送内容中的图片链接
	matches := re.FindAllStringSubmatch(text, -1)

	var contents []string
	for _, match := range matches {
		if len(match) >= 2 {
			contents = append(contents, match[2])
		}
	}

	return contents
}
