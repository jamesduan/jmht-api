package g

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

func CreateUser(user *User) *gorm.DB {
	db := DB.Create(user)
	return db
}

func generateToken() string {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	log.Println(token)
	return token
}

func GetToken(userID int) string {
	var token Token
	DB.Where(&Token{UserID: userID}).First(&token)
	return token.Token
}

func GetTokenByTokenString(token string) *Token {
	var wantToken Token
	DB.Where(&Token{Token: token}).First(&wantToken)
	return &wantToken
}

func DeleteToken(tokenstr string) bool {

	var token Token
	DB.Where("token = ?", tokenstr).First(&token)
	if token.ID == 0 {
		return false
	}
	db := DB.Delete(&token)
	if db.RowsAffected == 0 {
		return false
	} else {
		return true
	}

}

func CheckIsLogined(tokenstr string) bool {
	var token Token
	DB.Where("token = ?", tokenstr).First(&token)
	if token.ID == 0 {
		return false
	} else {
		return true
	}
}

func CheckLogin(username, password string, targetUser *User) error {
	// var user User
	var pass Password
	var tokenTmp Token
	log.Println(username, password)
	DB.Where(&User{Name: username}).First(targetUser)
	log.Println(targetUser.ID)

	if targetUser.ID != 0 {
		DB.Where(&Password{UserID: targetUser.ID}).First(&pass)
		targetUser.Password = &pass
		if targetUser.Password != nil {
			if targetUser.Password.PasswordMd5 == password {

				// DB.Where(&Token{UserID: targetUser.ID}).First(&tokenTmp)
				DB.Where(&Token{UserID: targetUser.ID}).First(&tokenTmp)

				if tokenTmp.ID == 0 {
					token := new(Token)
					token.UserID = targetUser.ID
					token.Token = generateToken()
					targetUser.Token = token
					DB.Save(targetUser)
					return nil
				} else {
					targetUser.Token = &tokenTmp
					return nil
				}
			} else {
				return fmt.Errorf("密码错误")
			}
		} else {
			return fmt.Errorf("用户密码还未设置")
		}
	}
	// log.Println(user.ID)
	return fmt.Errorf("用户不存在或者用户名错误")
}

func AddArticle(article *Article) bool {
	db := DB.Create(article)
	if db.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

func UpdateLikeById(id, like int) error {
	var article Article
	r1 := DB.First(&article, id)
	if r1.Error != nil {
		return fmt.Errorf("查不到此ID")
	}
	article.Like = like
	result := DB.Save(&article)
	if result.Error != nil {
		return fmt.Errorf("更新失败")
	}
	return nil
}

func UpdateArticle(id int, article *Article) error {
	var articleSave Article
	result := DB.First(&articleSave, id)
	if result.Error != nil {
		return fmt.Errorf("查不到")
	}
	articleSave.Title = article.Title
	articleSave.Content = article.Content
	result = DB.Save(&articleSave)
	if result.Error != nil {
		return fmt.Errorf("保存失败")
	}
	// log.Println(result)
	return nil
}

func DeleteArticle(id int) error {
	var article Article
	var image Image
	DB.First(&article, id)

	// 删除article和image
	tx := DB.Begin()

	db := DB.Delete(&article)

	if db.Error != nil {
		tx.Rollback()
		return db.Error
	}

	db = DB.Where(&Image{ArticleID: id}).First(&image)
	if db.Error != nil {
		tx.Rollback()
		return db.Error
	}

	db = DB.Delete(&image)
	if db.Error != nil {
		tx.Rollback()
		return db.Error
	}
	basedir := Config().Image.FilePath
	file := filepath.Join(basedir, image.FileName)
	err := os.Remove(file)
	if err != nil {
		tx.Rollback()
		return db.Error
	}
	tx.Commit()
	return nil
}

func Articles(articles *[]Article) (error, []Article) {
	query := DB.Order("id desc").Find(articles)
	if query.Error != nil {
		log.Println(query.Error)
		return query.Error, nil
	}
	var tmpArticles []Article
	for _, article := range *articles {
		// log.Println(article)
		err, img := getImageByArticleID(article.ID)
		if err != nil {
			log.Println(err)
			return err, nil
		}
		article.Image = img
		tmpArticles = append(tmpArticles, article)
	}

	return nil, tmpArticles
}

func getImageByArticleID(articleID int) (error, *Image) {
	var image Image
	result := DB.Where(&Image{ArticleID: articleID}).First(&image)
	// log.Println(image)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error, nil
	}
	return nil, &image
}

func getImageByProductID(productID int) (error, *Image) {
	var image Image
	result := DB.Where(&Image{ProductID: productID}).First(&image)
	log.Println(image)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error, nil
	}
	return nil, &image
}

func AddImage(image *Image) error {
	result := DB.Create(image)
	log.Println(result)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateCategory(category *ProductCategory) error {
	result := DB.Create(category)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func GetAllCategories(categories *[]ProductCategory) error {
	query := DB.Order("id desc").Find(categories)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func DeleteCategory(categoryID int) error {
	var category ProductCategory
	result := DB.First(&category, categoryID)
	if result.Error != nil {
		return result.Error
	}
	result = DB.Delete(&category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateProduct(pro *Product) error {
	var category ProductCategory
	result := DB.First(&category, pro.ProductCategoryID)
	if result.Error != nil {
		return result.Error
	}
	var set ProductSet
	result = DB.First(&set, pro.SetID)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	result = DB.Create(pro)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteProduct(id int) error {
	var pro Product
	result := DB.First(&pro, id)
	if result.Error != nil {
		return result.Error
	}
	result = DB.Delete(pro)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetSetByID(setID int) (*ProductSet, error) {
	var set ProductSet

	result := DB.First(&set, setID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &set, nil
}

func QueryProducts(products *[]Product) (error, []Product) {
	query := DB.Order("id desc").Find(products)
	if query.Error != nil {
		return query.Error, nil
	}
	var tmpProducts []Product

	log.Println(products)
	for _, product := range *products {
		// log.Println(article)
		err, img := getImageByProductID(product.ID)
		if err != nil {
			log.Println(err)
			product.Image = nil
			// return err, nil
		}
		product.Image = img
		log.Println(product.SetID)
		proSet, err1 := GetSetByID(product.SetID)
		if err1 != nil {
			log.Println(err1)
			product.ProductSet = nil
		}
		product.ProductSet = proSet

		tmpProducts = append(tmpProducts, product)
	}
	// log.Println(tmpProducts)
	return nil, tmpProducts
}

func ListProSet(sets *[]ProductSet) error {
	result := DB.Order("id desc").Find(sets)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func ProductSetCreate(set *ProductSet) error {
	result := DB.Create(set)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func ProductSetDel(setID int) error {
	var set ProductSet
	result := DB.First(&set, setID)
	if result.Error != nil {
		return result.Error
	}
	result = DB.Delete(&set)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
