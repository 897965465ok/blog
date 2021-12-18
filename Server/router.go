package main

import (
	controller "main/Controller"
	middleware "main/Middleware"

	"github.com/gin-gonic/gin"
)

// GET /product：列出所有商品
// POST /product：新建一个商品
// GET /product/ID：获取某个指定商品的信息
// PUT /product/ID：更新某个指定商品的信息
// DELETE /product/ID：删除某个商品
// GET /product/ID/purchase ：列出某个指定商品的所有投资者
// get /product/ID/purchase/ID：获取某个指定商品的指定投资者信息
func CollectRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.Cors())
	r.Static("/static", "./view/static")
	r.Static("/v1/markdown", "./markdown")
	r.LoadHTMLFiles("./view/index.tmpl")
	r.GET("/", controller.Index)
	r.GET("/ws", controller.WsCnection)
	r.NoRoute(controller.Index)
	V1 := r.Group("v1")
	{
		V1.GET("/picture", controller.BannerPiture)
		// 创建标签
		V1.POST("/tag", controller.CreateTag)
		// 查找全部标签
		V1.GET("/tags", controller.QueryAllTag)

		// 查询所有文章
		V1.GET("/articles", controller.QueryAllArticle)
		// 查询一个文章
		V1.GET("/article", controller.QueryOneArticle)
		V1.DELETE("/article", controller.DeleteArticle)
		// 创建文章
		V1.POST("/article", controller.CreateArticle)
		//  根据标签查询分类
		V1.GET("/query", middleware.AuthMiddleware(), controller.Query)
		// 发送邮件
		V1.GET("/email", controller.SendEamil)
		// 添加收藏网站
		V1.POST("/favorite", controller.ADDFavorite)
		//查询收藏网址
		V1.GET("/favorite", controller.QueryFavorites)
		// 删除收藏网址
		V1.DELETE("/favorite", controller.DeleteFavorite)
		// markdown
		V1.PUT("/markdowntohtml", controller.MarkdownToHmtl)

		//  注册
		V1.POST("/register", controller.Register)
		// 登陆
		V1.POST("/login", controller.Login)

		V1.GET("/watchnumber", controller.WatchNumber)
		V1.GET("/like", controller.Like)
		V1.POST("/comment", middleware.AuthMiddleware(), controller.Comment)
		/*以下都是测试接口 */
		// 人家服务器接口已废除
		V1.GET("/returnjson", controller.ReturnJson)
		//Wallhaven接口
		V1.GET("/wallhaven", controller.Wallhaven)
		// 支付宝测试接口
		V1.GET("/play", controller.Pay)
		// 支付宝回调接口
		V1.GET("/PayCallBack", controller.PayCallBack)
		// cookies
		V1.GET("/cookies", controller.Cookies)
		V1.GET("/openid", controller.Cookies)
		V1.GET("/comment", middleware.AuthMiddleware(), controller.GetComment)
		V1.POST("/wallhaven2", controller.Wallhaven_V2)
		V1.POST("/should", controller.TestShould)
		V1.GET("/oauth", controller.Oauth)

	}
	return r
}
