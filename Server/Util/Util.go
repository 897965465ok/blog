package util

import (
	"fmt"
	"io/ioutil"
	global "main/Global"
	model "main/Model"
	"math/rand"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

// type Response struct {
// 	Code    uint32      `json:"code"`
// 	Message uint32      `json:"message"`
// 	Data    interface{} `json:"data"`
// }

// type ResponseError struct {
// 	Code    uint32 `json:"code"`
// 	Message uint32 `json:"message"`
// }

var jwtKey = []byte(viper.GetString("jwt.key"))

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 生成token
func GeneractToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(720 * time.Minute)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), //token 有效时间
			IssuedAt:  time.Now().Unix(),     // // 发放时间
			Issuer:    "Jiang",               // 谁发放了这个token
			Subject:   "user token",          // 主题
			// Audience  string `json:"aud,omitempty"`
			// ExpiresAt int64  `json:"exp,omitempty"`
			// Id        string `json:"jti,omitempty"`
			// IssuedAt  int64  `json:"iat,omitempty"`
			// Issuer    string `json:"iss,omitempty"`
			// NotBefore int64  `json:"nbf,omitempty"`
			// Subject   string `json:"sub,omitempty"`
		},
	}
	//  生成jwt 三个部分
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 签名
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	clauns := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, clauns,
		func(token *jwt.Token) (i interface{}, err error) {
			return jwtKey, nil
		},
	)
	return token, clauns, err
}

func GetUserId(tokenString string) (uint, bool) {
	token, claims, err := ParseToken(tokenString)
	// token.Valid 如果令牌有效就是ture
	if token.Valid && err == nil {
		return claims.UserId, true
	} else {
		return 0, false
	}
}
func HasUserName(userName string) bool {
	DB := global.DB
	user := &model.User{}
	// 如果没有记录就等于0
	if DB.Where("name = ?", userName).First(user).RowsAffected == 0 {
		return true
	} else {
		return false
	}
}

// 加盐
func Encipherment(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	} else {
		return string(hash), nil
	}
}

// 解盐
func Decryption(hashPassword, password string) bool {
	// 正确密码验证
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}

//校验电子邮箱
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// 判断邮箱是否存在
func HasUserEmail(email string) bool {
	verify := VerifyEmailFormat(email)
	DB := global.DB
	user := &model.User{}
	// 如果没有记录就等于0
	if verify && (DB.Where("email = ?", email).First(user).RowsAffected == 0) {
		return true
	} else {
		return false
	}
}

// 发送邮箱
func SendEamil(mailTo []string, subject string, body string) error {

	//定义邮箱服务器连接信息，如果是网易邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"user": "897965465@qq.com",
		"pass": viper.GetString("pas"),
		"host": "smtp.qq.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int
	message := gomail.NewMessage()
	message.SetHeader("From", message.FormatAddress(mailConn["user"], "Jiang"))

	message.SetHeader("To", mailTo...)    //发送给多个用户
	message.SetHeader("Subject", subject) //设置邮件主题
	message.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err := d.DialAndSend(message)
	return err

}

// 响应
func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, message string) {
	ctx.JSON(httpStatus, gin.H{
		"code":    code,
		"data":    data,
		"message": message,
	})
}

// 成功
func Success(ctx *gin.Context, data gin.H, message string) {
	Response(ctx, http.StatusOK, http.StatusOK, data, message)
}

// 错误
func Fail(ctx *gin.Context, data gin.H, message string) {
	Response(ctx, http.StatusOK, http.StatusBadRequest, data, message)
}

//读取文件列表
func ReadFile(dirPath, dirName string) {
	DB := global.DB
	dirArr, err := ioutil.ReadDir(dirPath)
	ErrorHandling(err)
	for _, item := range dirArr {
		if item.IsDir() != true {
			fileName := item.Name()
			result := DB.Where("name = ?", fileName).First(&model.Article{})
			// 如果等于0就是不存在
			if result.RowsAffected == 0 {
				// input, _ := ioutil.ReadFile(filepath.Join(dirPath, fileName))
				// blackfriday.Run(input, blackfriday.WithNoExtensions())
				// localhost:3800/markdown/练习小项目/轮播图.md
				DB.Create(&model.Article{
					Name:        fileName,
					Tag:         dirName,
					Picture:     "",
					ArticlePath: fmt.Sprintf("http://localhost:3800/markdown/%s/%s", dirName, fileName),
				})
			}

		}

	}
	return
}

// 读取文件夹列表
func ReadDir(dirPath string) {
	DB := global.DB
	dirArr, err := ioutil.ReadDir(dirPath)
	ErrorHandling(err)
	for _, item := range dirArr {
		if item.IsDir() {
			dirName := item.Name()
			result := DB.Where("article_tag = ?", dirName).First(&model.Tags{})
			if result.RowsAffected == 0 {
				DB.Create(&model.Tags{ArticleTag: dirName})
			} else {
				dirPath, _ := filepath.Abs(filepath.Join(dirPath, dirName))
				ReadFile(dirPath, dirName)

			}

		}

	}
	return
}

//  错误处理
func ErrorHandling(err error) {
	if err != nil {
		panic(err)
	}
}

//生成count个[start,end)结束的不重复的随机数
func GenerateRangeNum(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}
