package model

type Adress struct {
	ID        int    `json:"id"`
	ShortName string `json:"shortname"`
	Address   string `json:"address"`
	User      User   `json:"user"`
}

type MobilePhone struct {
	ID          int    `json:"id"`
	ShortName   string `json:"shortname"`
	PhoneNumber string `json:"phone_number"`
	User        User   `json:"user"`
}

//User 用户模型用于登录，登出及授权的数据模型
type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Sex         string `json:"sex"`
	PortraitURL string `json:"portrait_url"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}
