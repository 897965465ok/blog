package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// 标签
type Tags struct {
	gorm.Model
	UUID       string `gorm:"type:varchar(50); not null;" json:"uuid"`
	ArticleTag string `gorm:"type:varchar(50);unique;not null;" json:"article_tag"`
}

// 保存文件路径

// 算个用户吧
type User struct {
	gorm.Model
	UUID     string `gorm:"type:varchar(50); not null;" json:"uuid"`
	Email    string `grom:"type:varchar(40); unique; not null;"  json:"email"`
	Name     string `grom:"type:varchar(40); not null;"  json:"name"`
	Password string `grom:"type:varchar(50); not null;"  json:"password"`
}

type Favorites struct {
	gorm.Model
	UUID string `gorm:"type:varchar(50); not null;" json:"uuid"`
	Name string `grom:"type:varchar(40); not null;"  json:"name"`
	Link string `grom:"type:varchar(300); not null;"  json:"link"`
}

type Banner struct {
	gorm.Model
	Name        string `grom:"type:varchar(40); not null;"  json:"name"`
	UUID        string `gorm:"type:varchar(50); not null;" json:"uuid"`
	Picture     string `gorm:"type:varchar(200);" json:"picture"`
	ArticlePath string `gorm:"type:varchar(255);"  json:"article_path" `
}
type Article struct {
	gorm.Model
	Picture      string `gorm:"type:varchar(200);" json:"picture"`
	Paragraph    string `gorm:"type:varchar(50);" json:"paragraph"`
	UUID         string `gorm:"type:varchar(50); not null;" json:"uuid"`
	Name         string `gorm:"type:varchar(70);" json:"name"`
	Tag          string `gorm:"type:varchar(50);"  json:"tag"`
	ArticlePath  string `gorm:"type:varchar(255);"  json:"article_path" `
	WhatchNumber int    `gorm:" DEFAULT 0 " json:"whatch_number"`
	Like         int    `gorm:" DEFAULT 0 " json:"like"`
	Hot          int    `gorm:" DEFAULT 0 " json:"hot"`
	REC          int    `gorm:" DEFAULT 0 " json:"rec"`
}

type Comments struct {
	gorm.Model
	UUID      string `gorm:"type:varchar(50); not null;" json:"uuid"`
	ArticleId string `gorm:"type:varchar(50); not null;" json:"Article_id"` // 文章id
	UserId    string `gorm:"type:varchar(50); not null;" json:"user_id"`
	Comment   string `gorm:"type:varchar(255); not null;" json:"comment"` // 内容
}
type TreePaths struct {
	Ancestor   string `gorm:"type:varchar(50); not null;" json:"ancestor"`
	Descendant string `gorm:"type:varchar(50); not null;" json:"descendant"`
	Distacne   string `gorm:"type:varchar(50); not null;" json:"distacne"`
}

type ImgUrl struct {
	gorm.Model
	Dimension_x int64  `json:"dimension_x"`
	Dimension_y int64  `json:"dimension_y"`
	File_size   int64  `json:"file_size"`
	UUID        string `gorm:"type:varchar(50);" json:"id"`
	Url         string `gorm:"type:varchar(255);" json:"url"`
	Short_url   string `gorm:"type:varchar(255);" json:"short_url"`
	Category    string `gorm:"type:varchar(50);" json:"category"`
	File_type   string `gorm:"type:varchar(50);" json:"file_type"`
	Path        string `gorm:"type:varchar(255);" json:"path"`
	Large       string `gorm:"type:varchar(255);" json:"large"`
	Original    string `gorm:"type:varchar(255);" json:"original"`
	Small       string `gorm:"type:varchar(255);" json:"small"`
}

func (ImgUrl *ImgUrl) BeforeCreate(tx *gorm.DB) error {
	ImgUrl.UUID = uuid.NewV4().String()
	return nil
}
func (Comments *Comments) BeforeCreate(tx *gorm.DB) error {
	Comments.UUID = uuid.NewV4().String()
	return nil
}

func (Favorites *Favorites) BeforeCreate(tx *gorm.DB) error {
	Favorites.UUID = uuid.NewV4().String()
	return nil
}

func (Article *Article) BeforeCreate(tx *gorm.DB) error {
	Article.UUID = uuid.NewV4().String()
	return nil
}

func (User *User) BeforeCreate(tx *gorm.DB) error {
	User.UUID = uuid.NewV4().String()
	return nil
}

func (Tags *Tags) BeforeCreate(tx *gorm.DB) error {
	Tags.UUID = uuid.NewV4().String()
	return nil
}

// type Author struct {
// 	Name string `gorm:"type:varchar(70);" json:"name"`
// 	// 作者头像
// 	Avatar string `gorm:"type:varchar(70);" json:"name"`
// 	// 作者id
// 	Id string `gorm:"type:varchar(50); not null;" json:"id"`
// 	// Status "normal"
// 	CommentID uint
// }

// type Comment struct {
// 	gorm.Model
// 	// 作者
// 	Author Author
// 	// 浏览量
// 	BrowseCount int `gorm:" DEFAULT 0 " json:"browse_count"`
// 	// 赞
// 	Like int `gorm:" DEFAULT 0 " json:"like"`
// 	// 图片
// 	Picture string `gorm:"type:varchar(200);" json:"picture"`
// 	// 段落
// 	Paragraph string `gorm:"type:varchar(50);" json:"paragraph"`
// 	// uuid和普通id一样
// 	UUID string `gorm:"type:varchar(50); not null;" json:"uuid"`
// 	// 文章名
// 	Name string `gorm:"type:varchar(70);" json:"name"`
// 	// 标签/分类
// 	Tag string `gorm:"type:varchar(50);"  json:"tag"`
// 	// 源文件所在路径
// 	ArticlePath string `gorm:"type:varchar(255);"  json:"article_path" `
// 	// 是否热门文章
// 	Hot int `gorm:" DEFAULT 0 " json:"hot"`

// 	REC int `gorm:" DEFAULT 0 " json:"rec"`
// }
