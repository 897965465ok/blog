package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
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
	UUID      string `gorm:"type:varchar(50); not null;" json:"uuid"`       // 主键
	ArticleId string `gorm:"type:varchar(50); not null;" json:"Article_id"` // 主键
	Comment   string `gorm:"type:varchar(255); not null;" json:"comment"`   // 主键
}
type TreePaths struct {
	Ancestor   int `gorm:" DEFAULT 0 " json:"ancestor"`
	Descendant int `gorm:" DEFAULT 0 " json:"descendant"`
	Distacne   int `gorm:" DEFAULT 0 " json:"distacne"`
}

func (Favorites *Favorites) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UUID", uuid.NewV4().String())
	return nil
}

func (Article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UUID", uuid.NewV4().String())
	return nil
}

func (User *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UUID", uuid.NewV4().String())
	return nil
}

func (Tags *Tags) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UUID", uuid.NewV4().String())
	return nil
}
