//Package gpbt200 狗屁不通文章生成器 调用api版
package gpbt200

import ( 
	"encoding/json"
	"io/ioutil"
	"net/http"

	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init(){
	engine := control.Register("gpbt200", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Help: "gpbt200\n" +
		"- 狗屁不通#关键字\n" +
		"例：狗屁不通#喝酸奶",
	})
	type gpbt struct {
		CODE int `json:"code"`
		TEXT string `json:"text"`
	}
	engine.OnRegex ("狗屁不通#(.*)").SetBlock(true).
	Handle(func(ctx *zero.Ctx){
		keyw := ctx.State["regex_matched"].([]string)[1]
		url := "http://ovooa.com/API/dog/api.php?msg="+ keyw +"&num=200&type=json"
		resp, _ := http.Get(url)
		s := gpbt{}
		body, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		json.Unmarshal([]byte(body), &s)
		ctx.SendChain(message.Text(s.TEXT))
	})
}
