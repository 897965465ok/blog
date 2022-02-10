package controller

import (
	"bytes"
	"fmt"
	"io/ioutil"
	global "main/Global"
	model "main/Model"
	util "main/Util"
	"net/http"
	"strconv"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func Oauth(ctx *gin.Context) {
	var clientId string
	var clientSecret string
	var Args string
	var Url string
	// origin := ctx.Request.Referer()
	// fmt.Println("域名", origin)
	code := ctx.Query("code")
	origin := ctx.Query("origin")
	if origin == "github" {
		fmt.Println("github访问")
		var redirect_uri string = "http://www.mrjiang.work/v1/oauth?origin=github"
		clientId = viper.GetString("GithubClientId")
		clientSecret = viper.GetString("GithubClientSecret")
		Url = "https://github.com/login/oauth/access_token?"
		Args = fmt.Sprintf("code=%s&client_id=%s&client_secret=%s&redirect_uri=%s", code, clientId, clientSecret, redirect_uri)

	} else {
		fmt.Println("gitee访问")
		var redirect_uri string = "http://www.mrjiang.work/v1/oauth?origin=gitee"
		clientId = viper.GetString("GieeClientId")
		clientSecret = viper.GetString("GieeClientSecret")
		Url = "https://gitee.com/oauth/token"
		Args = fmt.Sprintf("grant_type=authorization_code&code=%s&client_id=%s&redirect_uri=%s&client_secret=%s", code, clientId, redirect_uri, clientSecret)

	}
	response, err := http.Post(Url, "application/x-www-form-urlencoded", strings.NewReader(Args))
	if err != nil || response.StatusCode != http.StatusOK {
		ctx.Status(http.StatusServiceUnavailable) // 503
		fmt.Println("错误代码", response.StatusCode)
		return
	}
	defer response.Body.Close()
	str, _ := ioutil.ReadAll(response.Body)
	result := make(gin.H)
	result["code"] = 200
	jsonparser.ObjectEach(str, func(key, value []byte, dataType jsonparser.ValueType, offset int) error {
		result[string(key)] = string(value)
		return nil
	})

	ctx.HTML(http.StatusOK, "oauth", result)
}
func Index(Ginctx *gin.Context) {
	Ginctx.HTML(http.StatusOK, "index", gin.H{})
}

// 注册
func Register(ctx *gin.Context) {
	DB := global.DB
	user := &model.User{}
	userName := ctx.PostForm("userName")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")
	// 用户是否存在
	if DB.Where("email = ?", email).First(&user).Error == nil {
		util.Fail(ctx, nil, "用户已存在")
	} else {
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

	}

}

// 登陆
func Login(ctx *gin.Context) {
	DB := global.DB
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
	DB := global.DB
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
	DB := global.DB
	result := &[]model.Tags{}
	DB.Find(result)
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": result,
	})
}

// 创建文章
func CreateArticle(ctx *gin.Context) {
	DB := global.DB
	Article := &model.Article{Tag: ctx.Query("tag")}
	DB.Create(Article)
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": Article,
	})
}

// 查询全部文章

func QueryAllArticle(ctx *gin.Context) {
	DB := global.DB
	result := &[]model.Article{}
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	var count int64
	DB.Model(&model.Article{}).Count(&count)
	DB.Limit(limit).Offset(offset * limit).Find(result)
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": result,
		"count":  count,
	})
}

// 查询一个文章

func QueryOneArticle(ctx *gin.Context) {
	DB := global.DB
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
	DB := global.DB
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
	DB := global.DB
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
	DB := global.DB
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
	DB := global.DB
	result := &[]model.Favorites{}
	DB.Find(result)
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": result,
	})
}

// 删除一个
func DeleteFavorite(ctx *gin.Context) {
	DB := global.DB

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
	DB := global.DB
	UUID := ctx.Query("uuid")
	article := &model.Article{}
	DB.Where("uuid = ?", UUID).First(&article).Update("WhatchNumber", article.WhatchNumber+1)
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": article,
	})
}

// Like
func Like(ctx *gin.Context) {
	DB := global.DB
	UUID := ctx.Query("uuid")
	result := &model.Article{}
	DB.Where("uuid = ?", UUID).First(result).Update("Like", result.Like+1)
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": result,
	})
}

// 评论 文章
func Comment(ctx *gin.Context) {
	DB := global.DB
	// 判断是否回复文章
	tokenString := ctx.GetHeader("Authorization")[7:]
	userId, err := util.GetUserId(tokenString)
	if !err {
		ctx.JSON(http.StatusOK, gin.H{
			"code":   http.StatusBadRequest,
			"result": "失败",
		})
	}
	boolean, _ := strconv.ParseBool(ctx.PostForm("replyArticle"))
	uuid := ctx.PostForm("articleId")
	content := ctx.PostForm("content")
	user := model.User{}
	// inerface 转结构体
	// interUser, exists := ctx.Get("user")
	// //  user是否存在
	// if exists {
	// 	mapstructure.Decode(interUser, user)
	// }
	if DB.Where("id = ?", userId).First(&user).Error == nil {
		if !boolean {
			// 回复文章
			article := model.Article{UUID: uuid}
			if DB.Select("id").Where("uuid = ?", article.UUID).First(&article).Error == nil {
				// DB.Model(&user).Association("Comment").Append(&comment)
				// DB.Model(&article).Association("Comment").Append(&comment)
				DB.Create(&model.Comment{
					Content:   content,
					UserID:    user.ID,
					ArticleID: article.ID,
					Reply_Id:  article.ID,
					ISReply:   boolean,
				})
				ctx.JSON(http.StatusOK, gin.H{
					"code":   http.StatusOK,
					"result": "回复文章成功",
				})

			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"code":   http.StatusBadRequest,
					"result": "失败",
				})
			}
		} else {
			userComent := model.Comment{}
			if DB.Where("uuid = ?", uuid).First(&userComent).Error == nil {
				replyComent := model.Comment{
					Content: content,
					ISReply: boolean,
					To:      ctx.PostForm("userName"),
				}
				// DB.Model(&user).Association("Comment").Append(&replyComent)
				replyComent.UserID = user.ID
				userComent.Replys = append(userComent.Replys, replyComent)
				DB.Save(&userComent)
				ctx.JSON(http.StatusOK, gin.H{
					"code":   200,
					"result": "回复评论成功",
				})
			}

		}

	}

}

// 递归
func tree(commons []model.Comment) []map[string]interface{} {
	DB := global.DB
	Replys := []model.Comment{}
	var result []map[string]interface{}
	for _, item := range commons {
		user := model.User{}
		container := make(map[string]interface{})
		DB.Select("name", "uuid", "id", "email").Where("id=?", item.UserID).First(&user)
		DB.Where("manager_id=?", item.ID).Find(&Replys)
		// 结构体转map
		mapstructure.Decode(item, &container)
		container["User"] = user
		container["Replys"] = tree(Replys)
		result = append(result, container)
	}
	return result
}

// 获取评论
func GetComment(ctx *gin.Context) {
	DB := global.DB
	article := model.Article{UUID: ctx.Query("articleId")}
	commons := []model.Comment{}
	if DB.Select("id").Where("uuid=?", article.UUID).First(&article).Error == nil &&
		DB.Where("article_id=?", article.ID).Find(&commons).Error == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"result":  tree(commons),
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
