package controller

import (
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
	"github.com/panjf2000/ants/v2"
)

// markdown-to-html
func MarkdownToHmtl(ctx *gin.Context) {
	dir, _ := os.Getwd()
	util.ReadDir(filepath.Join(dir, "/markdown"))
	util.Success(ctx, nil, "成功")
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

// 获取轮播图
func GetBanner(ctx *gin.Context) {
	DB := common.GetDB()
	banner := []model.Banner{}
	if DB.Find(&banner).Error == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "成功",
			"result":  banner,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "失败",
			"result":  banner,
		})
	}
}

// 添加轮播图
func AppendBanner(ctx *gin.Context) {
	DB := common.GetDB()
	banner := []model.Banner{}
	for _, uuId := range ctx.PostForm("bannerIds") {
		DB.Raw(`SELECT 
		small,url,short_url,path,original,large,uuid as img_url_id 
		FROM img_url 
		WHERE uuid = ?`, uuId).
			Scan(&banner)
	}
	// 忽略字段
	if DB.Omit("id").Create(&banner).Error == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "成功",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    401,
			"message": "失败",
		})
	}
}

// 删除轮播图
func DeleteBanner(ctx *gin.Context) {

}

//  添加轮播图
func AddBanner(ctx *gin.Context) {

}

// 访问Wallhaven
// 访问Wallhaven
func Wallhaven(ctx *gin.Context) {
	var count int64
	DB := common.GetDB()
	Img := &[]model.ImgUrl{}
	DB.Model(&model.ImgUrl{}).Count(&count)
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	fmt.Println(limit * offset)
	err := DB.Select([]string{"small", "original", "large", "uuid", "id"}).Limit(limit).Offset(limit * offset).Find(Img).Error
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":   http.StatusNotFound,
			"result": err,
		})
	}
	container := make(map[string]interface{})
	container["count"] = count
	container["imgs"] = Img
	ctx.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"result": container,
	})
}
