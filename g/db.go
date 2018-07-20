package g

import (
	"log"
	"time"

	// _ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDBConfig() {
	var err error
	// DB, err = sql.Open("mysql", connectString)
	DB, err = gorm.Open("mysql", "root:redhat@tcp(127.0.0.1:3306)/jmht?loc=Local&parseTime=true")
	// log.Println(DB, err)
	if err != nil {
		log.Fatal("open db fail", err)
	}
	DB.DB().SetMaxOpenConns(100)
	DB.DB().SetMaxIdleConns(10)
	log.Println("Initialized Db Configuration.")
	// user
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Password{}, &Email{}, &Wechat{}, &Token{}, &Article{}, &Product{}, &Image{}, &ProductCategory{}, &ProductSet{})
	user := new(User)
	user.Name = "admin"
	user.NickName = "管理员"

	local, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation("2006-01-02", "1991-10-01", local)

	user.BirthDay = t
	user.RoleName = "admin"

	pass := new(Password)
	pass.PasswordMd5 = "888888"
	user.Password = pass

	DB.Create(user)
	// g.DB.Create(&g.User{Name: "jamesduan", NickName: "james"})
	// DB.Create(&User{Name: "jamesduan", NickName: "james"})
	// DB.Create(&model.Product{Code: "L1212", Price: 1000})
	// adresses := make([]model.HomeAddress, 0)
	// homeAddress := model.HomeAddress{Addess: "Shanghai", Post: "djkjflsdjflsdjlf"}
	// adresses = append(adresses, homeAddress)
	// DB.Create(&model.Users{Name: "jamesduan", Sex: "男", Number: "12345667", HomeAddresses: adresses, BirthDay: time.Unix(1469579899, 0)})
}
