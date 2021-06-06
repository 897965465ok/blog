package controller

import (
	"bytes"
	"fmt"
	"io/ioutil"
	common "main/Common"
	model "main/Model"
	util "main/Util"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	uuid "github.com/satori/go.uuid"
	"github.com/smartwalle/alipay/v3"
	"github.com/spf13/viper"
)

func Index(Ginctx *gin.Context) {
	Ginctx.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

// 注册
func Register(ctx *gin.Context) {
	DB := common.GetDB()
	user := &model.User{}
	DB.AutoMigrate(user)
	userName := ctx.PostForm("userName")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")
	if util.HasUserName(userName) {
		user.Name = userName
		user.Email = email
		// 判断加盐是否失败
		password, err := util.Encipherment([]byte(password))
		if err == nil {
			user.Password = password
		}
		DB.Create(user)
		util.Success(ctx, nil, "注册成功")
		return
	}
	util.Fail(ctx, nil, "用户已存在")
}

// 登陆
func Login(ctx *gin.Context) {
	DB := common.GetDB()
	user := &model.User{}
	userName := ctx.PostForm("userName")
	password := ctx.PostForm("password")

	if DB.Where("name = ?", userName).First(&user).Error == nil {
		fmt.Println(userName, password)
		if util.Decryption(user.Password, password) {
			token, _ := util.GeneractToken(*user)
			util.Success(ctx, gin.H{"token": token}, "登陆成功")
		}
	} else {
		util.Fail(ctx, nil, "登陆失败")
	}

}

// 获取轮播图URL
func BannerPiture(ctx *gin.Context) {

	response, err := http.Get("https://api.ixiaowai.cn/api/api.php")
	if err != nil || response.StatusCode != http.StatusOK {
		ctx.Status(http.StatusServiceUnavailable) // 503
		return
	}

	reader := response.Body
	contentLength := response.ContentLength
	contentType := response.Header.Get("contentType")

	// 作为消息主体中的消息头
	// 在 HTTP 场景中，第一个参数或者是 inline（默认值，表示回复中的消息体会以页面的一部分或者整个页面的形式展示），或者是 attachment（意味着消息体应该被下载到本地；大多数浏览器会呈现一个“保存为”的对话框，将 filename 的值预填为下载后的文件名，假如它存在的话）
	// Content-Disposition: inline
	// Content-Disposition: attachment
	// Content-Disposition: attachment; filename="filename.jpg"

	extraHeaders := map[string]string{
		"Content-Disposition": "inline",
	}

	ctx.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)

}

// 创建标签
func CreateTag(ctx *gin.Context) {
	DB := common.GetDB()
	DB.AutoMigrate(&model.Tags{})
	tags := &model.Tags{ArticleTag: ctx.Query("article_tag")}
	if DB.Create(tags).Error != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": "创建失败",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": tags,
	})
}

// 查询所有标签

func QueryAllTag(ctx *gin.Context) {
	DB := common.GetDB()
	result := DB.Find(&[]model.Tags{})
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": result.Value,
	})
}

// 创建文章
func CreateArticle(ctx *gin.Context) {
	DB := common.GetDB()
	DB.AutoMigrate(&model.Article{})
	Article := &model.Article{Tag: ctx.Query("tag")}
	result := DB.Create(Article)
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": result,
	})
}

// 查询全部文章

func QueryAllArticle(ctx *gin.Context) {
	DB := common.GetDB()
	result := DB.Find(&[]model.Article{})
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": result.Value,
	})
}

// 查询一个文章

func QueryOneArticle(ctx *gin.Context) {
	DB := common.GetDB()
	UUID := ctx.Query("uuid")
	result := DB.Where("uuid = ?", UUID).First(&model.Article{})
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": result,
	})
}

// DeleteArticle

func DeleteArticle(ctx *gin.Context) {
	DB := common.GetDB()
	UUID := ctx.Query("uuid")
	result := DB.Where("uuid = ?", UUID).First(&model.Article{})
	DB.Unscoped().Delete(result.Value)
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": result,
	})
}

// 根据标签查询分类
func Query(ctx *gin.Context) {
	DB := common.GetDB()
	var articles []model.Article
	Sql := fmt.Sprintf(
		`SELECT * FROM  article  WHERE article.tag = 
	(SELECT tags.article_tag FROM tags WHERE tags.article_tag ="%s");`,
		ctx.Query("value"))
	if DB.Raw(Sql).Scan(&articles).Error == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": articles,
		})
	}
}

// 发送邮件
func SendEamil(ctx *gin.Context) {
	//定义收件人
	mailTo := []string{"897965465@qq.com"}
	//邮件主题为"Hello"
	subject := "Jiang的博客"
	// 邮件正文
	body := "如果我是DJ你会爱我吗"

	if util.SendEamil(mailTo, subject, body) != nil {
		util.Fail(ctx, nil, "发送失败")
	} else {
		util.Success(ctx, nil, "发送成功")
	}

}

// markdown-to-html
func MarkdownToHmtl(ctx *gin.Context) {
	dir, _ := os.Getwd()
	util.ReadDir(filepath.Join(dir, "/markdown"))
	util.Success(ctx, nil, "成功")
}

// 返回格式为json
func ReturnJson(ctx *gin.Context) {
	response, err := http.Get("https://api.ixiaowai.cn/gqapi/gqapi.php?return=json")
	if err != nil || response.StatusCode != http.StatusOK {
		ctx.Status(http.StatusServiceUnavailable) // 503
		return
	}
	defer response.Body.Close()
	str, _ := ioutil.ReadAll(response.Body)
	text := bytes.TrimPrefix(str, []byte{239, 187, 191})
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"json": string(text),
	})

}

// 添加网址
func ADDFavorite(ctx *gin.Context) {
	DB := common.GetDB()
	DB.AutoMigrate(&model.Favorites{})
	Favorites := &model.Favorites{
		Name: ctx.Query("name"),
		Link: ctx.Query("link"),
	}
	result := DB.Create(Favorites)
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": result,
	})
}

// 查询全部收藏
func QueryFavorites(ctx *gin.Context) {
	DB := common.GetDB()
	result := DB.Find(&[]model.Favorites{})
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": result.Value,
	})
}

// 删除一个
func DeleteFavorite(ctx *gin.Context) {
	DB := common.GetDB()
	UUID := ctx.Query("uuid")
	result := &model.Favorites{}
	DB.Where("uuid = ?", UUID).First(result)
	DB.Unscoped().Delete(result)
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": result,
	})
}

// 观看次数
func WatchNumber(ctx *gin.Context) {
	DB := common.GetDB()
	UUID := ctx.Query("uuid")
	result := &model.Article{}
	DB.Where("uuid = ?", UUID).First(result)
	DB.First(result).Update("WhatchNumber", result.WhatchNumber+1)
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": result,
	})
}

// Like
func Like(ctx *gin.Context) {
	DB := common.GetDB()
	UUID := ctx.Query("uuid")
	result := &model.Article{}
	DB.AutoMigrate(&model.Article{})
	DB.Where("uuid = ?", UUID).First(result)
	DB.First(result).Update("Like", result.Like+1)
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": result,
	})
}

// 访问Wallhaven
func Wallhaven(ctx *gin.Context) {
	//创建任务列表
	ports := make(chan int, 9)
	// 任务结果
	results := make(chan string, cap(ports))
	for i := 0; i < cap(ports); i++ {
		// 开启线程
		go func(ports chan int, result chan string) {
			for page := range ports {
				Url := fmt.Sprintf("https://wallhaven.cc/api/v1/search?q=%s&page=%d", ctx.Query("q"), (page + 1))
				response, err := http.Get(Url)
				if err != nil || response.StatusCode != http.StatusOK {
					ctx.Status(http.StatusServiceUnavailable) // 503
					results <- strconv.Itoa(http.StatusServiceUnavailable)
					return
				}
				defer response.Body.Close()
				str, _ := ioutil.ReadAll(response.Body)
				value, _, _, _ := jsonparser.Get(str)
				data, err := jsonparser.GetUnsafeString(value, "data")
				if err != nil {
					fmt.Println("获取data失败")
				}
				results <- data

			}
		}(ports, results)
	}
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 0
	}
	go func() {
		// 派发任务
		for i := 0; i < 9; i++ {
			ports <- page + i
		}
		close(ports)
	}()
	array := make([]string, 0)
	for i := 0; i < cap(results); i++ {
		data, ok := <-results
		if ok == false {
			close(results)
			break
		}
		array = append(array, data)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": array,
	})

}

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

// 评论 文章
func Comment(ctx *gin.Context) {
	DB := common.GetDB()
	articleId := ctx.PostForm("articleId")
	content := ctx.PostForm("content")
	DB.Set("gorm:table_options", "ENGINE=InnoDB CHARACTER SET utf8mb4").AutoMigrate(&model.Comments{}, &model.TreePaths{})
	comment := &model.Comments{}
	treePaths := &model.TreePaths{}
	article := &model.Article{}
	user := &model.User{}
	interUser, _ := ctx.Get("user")
	mapstructure.Decode(interUser, user)
	DB.Where("uuid = ?", articleId).First(article)
	comment.ArticleId = article.UUID
	comment.Comment = content
	comment.UserId = user.UUID
	treePaths.Ancestor = article.UUID
	DB.Create(comment)
	treePaths.Descendant = comment.UUID
	DB.Create(treePaths)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "成功",
	})
}

// 获取评论
func GetComment(ctx *gin.Context) {
	type Result struct {
		UserName string
		UserID   string
		Content  string
	}
	articleId := ctx.Query("articleId")
	DB := common.GetDB()
	comments := &[]model.Comments{}
	user := &model.User{}
	// 查出所有评论
	DB.Where("article_id  = ?", articleId).Find(comments)
	results := make([]Result, len(*comments), len(*comments))

	for index, value := range *comments {
		DB.Where("uuid = ?", value.UserId).First(user)
		results[index].Content = value.Comment
		results[index].UserName = user.Name
		results[index].UserID = user.UUID
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": results,
	})

}
