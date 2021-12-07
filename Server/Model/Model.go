package model

import (
	"database/sql/driver"
	"strings"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// 标签
type Tags struct {
	gorm.Model
	UUID       string `gorm:"type:varchar(50); not null;" json:"uuid"`
	ArticleTag string `gorm:"type:varchar(50);unique;not null;" json:"article_tag"`
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

// 算个用户吧
type User struct {
	gorm.Model
	UUID string `gorm:"type:varchar(50); not null;" json:"uuid"`
	// 作者头像
	Avatar string
	// 工作类型 前端工程师
	Professional string
	// 状态
	Status      string
	Email       string    `grom:"type:varchar(40); unique; not null;"  json:"email" `
	Name        string    `grom:"type:varchar(40); not null;"  json:"name"`
	Password    string    `grom:"type:varchar(50); not null;"  json:"password;"`
	Comment     []Comment `gorm:"polymorphic:User;"`
	ArticleID   uint
	ArticleType string
}

type Comment struct {
	gorm.Model
	UserID   uint
	UserType string

	User User

	ArticleID   uint
	ArticleType string
	UUID        string `gorm:"type:varchar(50); not null;" json:"uuid"`
	// 暂时不做
	// Replys      []Comment `gorm:"foreignkey:ManagerID"`
	// ManagerID *uint
	// 回复内容
	Content string
	// 回复用户还是文章
	ISReply bool
	// 回复用的ID
	Reply_Id uint
	// 回复谁
	// To string
}

// 文章
type Article struct {
	gorm.Model
	User    *User     `gorm:"polymorphic:Article;"`
	Comment []Comment `gorm:"polymorphic:Article;"`
	UUID    string    `gorm:"type:varchar(50); not null;" json:"uuid"`
	// 浏览量
	// 	BrowseCount int64
	WhatchNumber int `gorm:" DEFAULT 0 " json:"whatch_number"`
	// 	// 类型
	// 	Classify string
	Tag string `gorm:"type:varchar(50);"  json:"tag"`
	// 文章内的图片 自定义类型
	Cover     Args
	Picture   string `gorm:"type:varchar(200);" json:"picture"`
	Paragraph string `gorm:"type:varchar(50);" json:"paragraph"`
	// 标题
	Name string `gorm:"type:varchar(70);" json:"name"`

	ArticlePath string `gorm:"type:varchar(255);"  json:"article_path" `
	// 	// 收集次数
	// 	CollectionCount int64
	Like int `gorm:" DEFAULT 0 " json:"like"`
	Hot  int `gorm:" DEFAULT 0 " json:"hot"`
	REC  int `gorm:" DEFAULT 0 " json:"rec"`

	// 评论数
	CommentsCount int64
	// 文章内容现在未
	Content string
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

type Args []string

// 实现 自定义类型
func (str Args) Value() (driver.Value, error) {
	if len(str) > 0 {
		var s string = str[0]
		for _, v := range str[1:] {
			s += "," + v
		}

		return s, nil
	}
	return "", nil
}

func (str *Args) Scan(value interface{}) error {
	bytes, _ := value.([]byte)
	*str = strings.Split(string(bytes), ",")
	return nil
}

func (ImgUrl *ImgUrl) BeforeCreate(tx *gorm.DB) error {
	ImgUrl.UUID = uuid.NewV4().String()
	return nil
}
func (Comment *Comment) BeforeCreate(tx *gorm.DB) error {
	Comment.UUID = uuid.NewV4().String()
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

// type Base struct {
// 	ID        uuid.UUID  `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
// 	CreatedAt time.Time  `json:"created_at"`
// 	UpdatedAt time.Time  `json:"updated_at"`
// 	DeletedAt *time.Time `gorm:"index"`
// }

// func (Base *Base) BeforeCreate(tx *gorm.DB) (err error) {
// 	// UUID version 4
// 	Base.ID = uuid.New()
// 	return
// }

// type Author struct {
// 	gorm.Model
// 	// 作者名字
// 	AuthorName string
// 	// 作者头像
// 	Avatar string
// 	// 状态
// 	Status string
// 	// 工作类型 前端工程师
// 	Professional string
// 	OwnerID      int
// 	OwnerType    string
// }

// type Comment struct {
// 	gorm.Model
// 	// 评论这信息
// 	Author Author `gorm:"polymorphic:Owner;"`
// 	// 回复的用户
// 	Replys []Comment `gorm:"polymorphic:Comment;"`
// 	// 多态引用
// 	CommentID   int
// 	CommentType string
// 	// 回复内容
// 	Content string
// 	// 回复用户还是文章
// 	ISReply bool
// 	// 回复用的ID
// 	Reply_Id uint
// 	// 回复谁
// 	To string
// }

// type Article struct {
// 	gorm.Model
// 	// 作者信息
// 	Author Author `gorm:"polymorphic:Owner;"`
// 	// 评论人列表
// 	Comment []Comment `gorm:"polymorphic:Comment;"`
// 	// 浏览量
// 	BrowseCount int64
// 	// 类型
// 	Classify string
// 	// 收集次数
// 	CollectionCount int64
// 	// 评论数
// 	CommentsCount int64
// 	// 文章内容
// 	Content string
// 	// 文章包含图片
// 	Cover Args
// 	// 模式 column
// 	Mode string
// 	// 点赞 次数
// 	ThumbsUpCount int64
// 	// Title
// 	Title string
// }
