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
	"sync"

	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"github.com/panjf2000/ants/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/smartwalle/alipay/v3"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func Index(Ginctx *gin.Context) {
	Ginctx.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

// 注册
func Register(ctx *gin.Context) {
	DB := common.GetDB()
	sqlDB, _ := DB.DB()
	defer sqlDB.Close()
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
	sqlDB, _ := DB.DB()
	defer sqlDB.Close()
	user := &model.User{}
	userName := ctx.PostForm("userName")
	password := ctx.PostForm("password")

	if DB.Where("name = ?", userName).First(&user).Error == nil {
		fmt.Println(userName, password)
		if util.Decryption(user.Password, password) {
			token, _ := util.GeneractToken(*user)
			util.Success(ctx, gin.H{"token": token}, "登陆成功")
		} else {
			util.Fail(ctx, nil, "请用户重新登录")
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
	sqlDB, _ := DB.DB()
	defer sqlDB.Close()
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
	sqlDB, _ := DB.DB()
	defer sqlDB.Close()
	result := &[]model.Tags{}
	DB.Find(result)
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": result,
	})
}

// 创建文章
func CreateArticle(ctx *gin.Context) {
	DB := common.GetDB()
	sqlDB, _ := DB.DB()
	defer sqlDB.Close()
	DB.AutoMigrate(&model.Article{})
	Article := &model.Article{Tag: ctx.Query("tag")}
	DB.Create(Article)
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": Article,
	})
}

// 查询全部文章

func QueryAllArticle(ctx *gin.Context) {
	DB := common.GetDB()
	sqlDB, _ := DB.DB()
	defer sqlDB.Close()
	result := &[]model.Article{}
	DB.Find(result)
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": result,
	})
}

// 查询一个文章

func QueryOneArticle(ctx *gin.Context) {
	DB := common.GetDB()
	sqlDB, _ := DB.DB()
	defer sqlDB.Close()
	UUID := ctx.Query("uuid")
	result := &model.Article{}
	DB.Where("uuid = ?", UUID).First(result)
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": result,
	})
}

// DeleteArticle

func DeleteArticle(ctx *gin.Context) {
	DB := common.GetDB()
	sqlDB, _ := DB.DB()
	defer sqlDB.Close()
	UUID := ctx.Query("uuid")
	result := &model.Article{}
	DB.Where("uuid = ?", UUID).First(result)
	DB.Unscoped().Delete(result)
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": result,
	})
}

// 根据标签查询分类
func Query(ctx *gin.Context) {
	DB := common.GetDB()
	sqlDB, _ := DB.DB()
	defer sqlDB.Close()
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
	sqlDB, _ := DB.DB()
	defer sqlDB.Close()
	DB.AutoMigrate(&model.Favorites{})
	Favorites := &model.Favorites{
		Name: ctx.Query("name"),
		Link: ctx.Query("link"),
	}
	DB.Create(Favorites)
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": Favorites,
	})
}

// 查询全部收藏
func QueryFavorites(ctx *gin.Context) {
	DB := common.GetDB()
	sqlDB, _ := DB.DB()
	defer sqlDB.Close()
	result := &[]model.Favorites{}
	DB.Find(result)
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": result,
	})
}

// 删除一个
func DeleteFavorite(ctx *gin.Context) {
	DB := common.GetDB()
	sqlDB, _ := DB.DB()
	defer sqlDB.Close()
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
	sqlDB, _ := DB.DB()
	defer sqlDB.Close()
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
	sqlDB, _ := DB.DB()
	defer sqlDB.Close()
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
	sqlDB, _ := DB.DB()
	// 创建数据库
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{}, &model.Comment{}, &model.Article{})
	defer sqlDB.Close()
	article := model.Article{
		UUID: ctx.PostForm("articleId"),
	}
	comment := model.Comment{
		Content: ctx.PostForm("content"),
	}

	user := model.User{}
	// inerface 转结构体
	interUser, exists := ctx.Get("user")
	//  user是否存在
	if exists {
		mapstructure.Decode(interUser, user)
	}
	if DB.First(&user).Error == nil && DB.Where("uuid = ?", article.UUID).First(&article).Error == nil {
		// 保存回复谁的ID
		comment.Reply_Id = article.ID
		// 文章还是回复评论
		comment.ISReply = false
		comment.User = []*model.User{&user}
		article.Comment = []*model.Comment{&comment}
		// 创造模型
		DB.Save(&article)
		util.Success(ctx, gin.H{
			"code":   http.StatusOK,
			"result": article}, "成功")
	} else {
		util.Fail(ctx, gin.H{
			"code":   http.StatusBadRequest,
			"result": "error",
		}, "服务器抽风了~~")
	}

}

// 获取评论
func GetComment(ctx *gin.Context) {
	DB := common.GetDB()
	sqlDB, _ := DB.DB()
	defer sqlDB.Close()
	article := model.Article{
		UUID: ctx.Query("articleId"),
	}
	if DB.Select("id").
		Where("uuid=?", article.UUID).
		First(&article).Error == nil &&
		DB.Model(&article).
			Preload("Comment.User",
				func(DB *gorm.DB) *gorm.DB {
					return DB.Select("ID", "name")
				}).
			Preload("Comment").
			First(&article).Error == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"result":  article.Comment,
			"message": "成功",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "error",
			"message": "失败",
		})
	}
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Child struct {
	Hub        *Hub
	Message    chan []byte
	UUID       string
	Conn       *websocket.Conn
	IsRegister bool
}

type Hub struct {
	Children   []*Child
	Broadcast  chan []byte
	Register   chan *Child
	Unregister chan *Child
}

var hub *Hub = &Hub{
	Children:   []*Child{},
	Broadcast:  make(chan []byte),
	Register:   make(chan *Child),
	Unregister: make(chan *Child),
}

func ReadLoop(c *Child) {
	defer func() {
		c.Conn.Close()
		for i, item := range c.Hub.Children {
			if item.UUID == c.UUID {
				c.Hub.Children = append(c.Hub.Children[0:i], c.Hub.Children[(i+1):]...)
				fmt.Println(c.Hub.Children)
			}
		}

	}()
	var (
		message []byte
		err     error
	)
	for {
		if _, message, err = c.Conn.ReadMessage(); err != nil {
			// 关闭通道
			return
		}
		c.Hub.Broadcast <- message
	}

}

func WriteLoop(c *Child) {
	defer func() {
		c.Conn.Close()
		for i, item := range c.Hub.Children {
			if item.UUID == c.UUID {
				c.Hub.Children = append(c.Hub.Children[0:i], c.Hub.Children[(i+1):]...)
				fmt.Println(c.Hub.Children)
			}
		}
	}()
	for {
		select {
		case message, err := <-c.Message:
			if err {
				c.Conn.WriteMessage(websocket.TextMessage, message)
			}
		}
	}

}
func WsCnection(ctx *gin.Context) {
	var (
		conn *websocket.Conn
		err  error
	)
	go func() {
		for {
			select {
			case child := <-hub.Register:
				//标记为已注册
				child.IsRegister = true
				// 存入孩子列表
				hub.Children = append(hub.Children, child)
			case message := <-hub.Broadcast:
				for _, item := range hub.Children {
					if item.IsRegister {
						item.Message <- message
					}
				}
			}
		}

	}()
	// 升级为ws协议
	if conn, err = upGrader.Upgrade(ctx.Writer, ctx.Request, nil); err != nil {
		return
	}
	c := &Child{
		Hub:        hub,
		Conn:       conn,
		UUID:       uuid.NewV4().String(),
		IsRegister: false,
		Message:    make(chan []byte, 256),
	}
	c.Hub.Register <- c
	go ReadLoop(c)
	go WriteLoop(c)

}

// 访问Wallhaven
func Wallhaven(ctx *gin.Context) {
	DB := common.GetDB()
	sqlDB, _ := DB.DB()
	defer sqlDB.Close()
	Img := &[]model.ImgUrl{}
	err := DB.Limit(216).Offset(1).Select([]string{"small", "original", "large", "uuid", "id"}).Find(Img).Error
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":   http.StatusNotFound,
			"result": err,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"result": Img,
	})
}
func Wallhaven_V2(ctx *gin.Context) {
	var wg sync.WaitGroup
	// defer ants.Release()
	pool, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		wg.Done()
		DB := common.GetDB()
		sqlDB, _ := DB.DB()
		defer sqlDB.Close()
		i, _ = i.(int64)
		Url := fmt.Sprintf("https://wallhaven.cc/api/v1/search?q=%s&page=%d", "anime", i)
		response, err := http.Get(Url)
		if err != nil || response.StatusCode != http.StatusOK {
			ctx.Status(http.StatusServiceUnavailable) // 503
			return
		}
		defer response.Body.Close()
		str, _ := ioutil.ReadAll(response.Body)
		jsonparser.ArrayEach(str, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			Img := &model.ImgUrl{}
			DB.AutoMigrate(Img)
			jsonparser.ObjectEach(value, func(key, value []byte, dataType jsonparser.ValueType, offset int) error {
				Strkey := string(key)
				StrValue := string(value)
				switch Strkey {
				case "url":
					Img.Url = StrValue
				case "short_url":
					Img.Short_url = StrValue
				case "category":
					Img.Category = StrValue
				case "dimension_x":
					v, _ := strconv.Atoi(StrValue)
					Img.Dimension_x = int64(v)
				case "dimension_y":
					v, _ := strconv.Atoi(StrValue)
					Img.Dimension_y = int64(v)
				case "file_size":
					v, _ := strconv.Atoi(StrValue)
					Img.File_size = int64(v)
				case "file_type":
					Img.File_type = StrValue
				case "path":
					Img.Path = StrValue
				case "thumbs":
					jsonparser.ObjectEach(value, func(key, value []byte, dataType jsonparser.ValueType, offset int) error {
						switch string(key) {
						case "large":
							Img.Large = string(value)
						case "original":
							Img.Original = string(value)
						case "small":
							Img.Small = string(value)
						}
						return nil
					})
				}
				return nil
			})
			DB.Create(Img)
		}, "data")
		if err != nil {
			fmt.Println("获取data失败", err)
		}
	})
	defer pool.Release()
	// 逐个提交任务。
	taskNumber := util.GenerateRangeNum(1, 200, 10)
	for _, item := range taskNumber {
		wg.Add(1)
		_ = pool.Invoke(int64(item))
	}
	wg.Wait()
	ctx.JSON(http.StatusOK, gin.H{
		"code":      200,
		"img_index": taskNumber,
	})
}
