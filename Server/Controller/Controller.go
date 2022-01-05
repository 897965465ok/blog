package controller

import (
	"fmt"
	util "main/Util"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/smartwalle/alipay/v3"
	"github.com/spf13/viper"
)

// 支付测试接口
func Pay(ctx *gin.Context) {
	var privateKey string = viper.GetString("privateKey")
	var publieKey string = viper.GetString("publieKey")
	var appId string = viper.GetString("appid")
	// client.LoadAppPublicCertFromFile("appCertPublicKey_2017011104995404.crt") // 加载应用公钥证书
	// client.LoadAliPayRootCertFromFile("alipayRootCert.crt") // 加载支付宝根证书
	// client.LoadAliPayPublicCertFromFile("alipayCertPublicKey_RSA2.crt") // 加载支付宝公钥证书
	// 生成环境要换成true
	var client, err = alipay.New(appId, privateKey, false)
	if err != nil {
		panic(err)
	}
	client.LoadAliPayPublicKey(publieKey)
	url, err := client.TradePagePay(alipay.TradePagePay{
		Trade: alipay.Trade{
			Subject:     ctx.Query("subject"),     // 订单标题
			OutTradeNo:  uuid.NewV4().String(),    // 商户订单号，64个字符以内、可包含字母、数字、下划线；需保证在商户端不重复
			TotalAmount: "0.1",                    // 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
			ProductCode: "FAST_INSTANT_TRADE_PAY", // 销售产品码，与支付宝签约的产品码名称。 注：目前仅支持FAST_INSTANT_TRADE_PAY
			NotifyURL:   "http://146.56.206.160/v1/PayCallBack",
			ReturnURL:   "http://146.56.206.160/v1/PayCallBack",
		}})
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": url.String(),
		// "param":  alipay.Params(),
	})

}

// 支付回调
func PayCallBack(ctx *gin.Context) {
	ctx.Request.ParseForm()
	for k, v := range ctx.Request.PostForm {
		fmt.Printf("k:%v\n", k)
		fmt.Printf("v:%v\n", v)
	}

}

//  cookies 测试cookies
func Cookies(ctx *gin.Context) {
	ctx.SetCookie("gay", "you are gay", 10, "/", "http://loclhost", false, false)
	value, _ := ctx.Cookie("gay")
	ctx.JSON(http.StatusOK, gin.H{
		"Cookie": value,
	})
}

// openid
func Openid(ctx *gin.Context) {
	ctx.Query("appid")

}
func TestShould(ctx *gin.Context) {
	util.Success(ctx, gin.H{
		"code":   http.StatusOK,
		"result": "测试函数"}, "成功")

}
