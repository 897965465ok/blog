package controller

import (
	"fmt"
	"io/ioutil"
	global "main/Global"
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

// @BasePath /api/v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// markdown-to-html
func MarkdownToHmtl(ctx *gin.Context) {
	dir, _ := os.Getwd()
	util.ReadDir(filepath.Join(dir, "/markdown"))
	util.Success(ctx, nil, "成功")
}

// @Tags admin
// @title Wallhaven_V2
// @version 1.0
// @Accept       json
// @Produce      json
// @license.name Apache 2.0
// @BasePath /admin/wallhaven_V2
// @Router /admin/wallhaven_V2 [post]
func Wallhaven_V2(ctx *gin.Context) {
	var wg sync.WaitGroup
	// defer ants.Release()
	pool, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		wg.Done()
		DB := global.DB
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

// @Tags admin
// @title getBanner
// @summary 获取轮播图
// @version 1.0
// @Accept       json
// @Produce      json
// @Param        limit    query     number  true  "数量"
// @Param        offset   query     number  true  "第几页"
// @license.name Apache 2.0
// @BasePath /admin/getBanner
// @Response 200 {object} ResponseError
// @Router /admin/getBanner [post]
func GetBanner(ctx *gin.Context) {
	DB := global.DB
	var count int64
	banner := []model.Banner{}
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	DB.Model(&model.Banner{}).Count(&count)
	if DB.Limit(limit).Offset(limit*offset).Find(&banner).Error == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "成功",
			"result":  banner,
			"count":   count,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "失败",
			"result":  banner,
			"count":   count,
		})
	}
}

// 添加轮播图
func AppendBanner(ctx *gin.Context) {
	DB := global.DB
	banner := []model.Banner{}
	temp := model.Banner{}
	jsonparser.ArrayEach([]byte(ctx.PostForm("bannerIds")), func(uuid []byte, dataType jsonparser.ValueType, offset int, err error) {
		DB.Raw(`SELECT
		small,url,short_url,path,original,large,uuid as img_url_id
		FROM img_url
		WHERE uuid = ?`, string(uuid)).Scan(&temp)
		banner = append(banner, temp)
	})
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
	DB := global.DB
	id := ctx.Query("id")
	SQL := fmt.Sprintf(`DELETE FROM banner WHERE id=%s LIMIT 1`, id)
	if DB.Exec(SQL).Error == nil {
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

// 访问Wallhaven
func Wallhaven(ctx *gin.Context) {
	var count int64
	DB := global.DB
	Imgs := make([]map[string]interface{}, 0)
	DB.Model(&model.ImgUrl{}).Count(&count)
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	err := DB.Model(&model.ImgUrl{}).
		Select([]string{"small", "original", "large", "uuid", "id"}).
		Limit(limit).Offset(limit * offset).Find(&Imgs).Error
	for _, item := range Imgs {
		// 判断是否存在
		if RowsAffected := DB.Where("img_url_id=?", item["uuid"]).
			Find(&model.Banner{}).
			RowsAffected; RowsAffected > 0 {
			item["check"] = true
		}
	}
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":   http.StatusNotFound,
			"result": err,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"result": map[string]interface{}{
			"count": count,
			"imgs":  Imgs,
		},
	})
}

func QueryUser(ctx *gin.Context) {
	DB := global.DB
	result := &[]model.User{}
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	offset = limit * offset
	// DB.Limit(limit).Offset(offset * limit).Find(result)
	sql := fmt.Sprintf(`select uuid,name,email from user limit %d,%d`, offset, limit)
	DB.Raw(sql).Scan(result)
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": result,
	})
}
