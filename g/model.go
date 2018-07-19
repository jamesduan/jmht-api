package g

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:255;unique"`
	NickName string
	Age      int
	BirthDay time.Time
	Emails   []Email
	Password *Password
	Wechat   *Wechat

	Token    *Token
	RoleName string
}

type Token struct {
	gorm.Model
	Token  string
	UserID int
}

type Email struct {
	gorm.Model
	UserID int
	Email  string
}

type Password struct {
	gorm.Model
	PasswordMd5  string
	PasswordSha1 string
	UserID       int
}

type Wechat struct {
	gorm.Model
	Name     string
	NickName string
	avator   string
	UserID   int
}

type Article struct {
	gorm.Model
	Title   string
	Content string `gorm:"size:4096"`
	Comment string `gorm:"size:1024"`
	Like    int    `gorm:"default:0"`
	UserID  int
	Image   *Image
}

type Image struct {
	gorm.Model

	UserID    int
	ArticleID int
	ProductID int

	SHA1     string
	FileName string
}

type Product struct {
	gorm.Model

	Name              string
	Description       string
	UserID            int
	Image             *Image
	ProductCategoryID int
}

type ProductCategory struct {
	gorm.Model

	Name string
}

func (User) TableName() string {
	return "user"
}

func (Email) TableName() string {
	return "u_email"
}

func (Password) TableName() string {
	return "u_password"
}

func (Wechat) TableName() string {
	return "u_wechat"
}

func (Token) TableName() string {
	return "u_token"
}

func (Article) TableName() string {
	return "article"
}

func (Product) TableName() string {
	return "product"
}

func (Image) TableName() string {
	return "image"
}

func (ProductCategory) TableName() string {
	return "product_category"
}
