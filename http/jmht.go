package http

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"jmht-api/g"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UploadImage struct {
	Image     string `json:"file"`
	ImageType string `json:"file_type"`
	ArticleID int    `json:"article_id"`
	ProductID int    `json:"product_id"`
}

type ResponseStatus struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func getSHA1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func decodeBase64(encodeString string) ([]byte, error) {
	b64data := encodeString[strings.IndexByte(encodeString, ',')+1:]
	decodeBytes, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		return nil, err
	}
	return decodeBytes, nil
}

func setCookies(w http.ResponseWriter, arr []*http.Cookie) {
	for _, cookie := range arr {
		http.SetCookie(w, cookie)
	}
}

func getCookie(k, v string) *http.Cookie {
	tNow := time.Now()
	return &http.Cookie{Name: k, Value: v, Expires: tNow.AddDate(1, 0, 0), Path: "/"}
}

func CheckLoginStatus(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// log.Println(r.URL.Path)
		// log.Println(r.BasicAuth)
		// log.Println(r.Cookies())

		// cookie, err := r.Cookie("user_id")
		// if err != nil {
		// 	log.Println(err)
		// }
		// log.Println(cookie.Name, cookie.Value)
		// userID, err := strconv.Atoi(cookie.Value)
		// if err != nil {
		// 	log.Println(err)
		// }
		// if ok := g.CheckIsLogined(userID); !ok {
		// 	RenderDataJson(w, map[string]ResponseStatus{"auth": ResponseStatus{Status: 0, Message: "还未登录"}})
		// } else {
		// 	f(w, r)
		// }
		if ok := g.CheckIsLogined(getHeaderToken(r)); !ok {
			RenderDataJson(w, map[string]ResponseStatus{"auth": ResponseStatus{Status: 0, Message: "还未登录"}})
		} else {
			f(w, r)
		}
	}
}

func ParsePost(model interface{}, r *http.Request) error {
	r.ParseForm()
	// fmt.Println(r.Form)
	body, _ := ioutil.ReadAll(r.Body)
	// user := new(User)
	if err := json.Unmarshal(body, model); err != nil {
		// fmt.Println(err)
		// RenderDataJson(w, map[string]string{"errorMsg": err.Error()})
		return err
	}
	return nil
}

func getUserIDFromCookie(r *http.Request) int {
	cookie, err := r.Cookie("user_id")
	if err != nil {
		log.Println(err)
	}
	// log.Println(cookie.Name, cookie.Value)
	userID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Println(err)
	}
	return userID
}

func getHeaderToken(r *http.Request) string {
	return r.Header.Get("token")
}

type ArticleRequest struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Comment string `json:"comment"`
	Like    int    `json:"like"`
	UserID  int    `json:"user_id"`
}

type ProductRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryID  int    `json:"category_id"`
}

type CategoryRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getTime(timestr string) time.Time {
	local, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation("2006-01-02", timestr, local)
	return t
}

func uploadImg(r *http.Request) error {
	var baseDir string = g.Config().Image.FilePath
	r.ParseForm()
	// fmt.Println(r.Form)
	body, _ := ioutil.ReadAll(r.Body)
	// log.Println(body)
	// user := new(User)
	image := new(UploadImage)
	if err := json.Unmarshal(body, image); err != nil {
		fmt.Println(err)
		return err
		// RenderDataJson(w, map[string]string{"errorMsg": err.Error()})
	}
	var img = new(g.Image)
	img.SHA1 = getSHA1(image.Image)
	img.ArticleID = image.ArticleID
	img.ProductID = image.ProductID
	token := getHeaderToken(r)
	userID := g.GetTokenByTokenString(token).UserID
	img.UserID = userID
	// log.Println(image.Image)
	decodedBytes, err := decodeBase64(image.Image)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("init base dir: ", baseDir)
	filePath := filepath.Join(baseDir, img.SHA1+"."+image.ImageType)
	log.Println(filePath)

	err = ioutil.WriteFile(filePath, decodedBytes, 0666)
	if err != nil {
		return err
	}
	img.FileName = img.SHA1 + "." + image.ImageType
	err = g.AddImage(img)
	if err != nil {
		return err
	}
	return nil
}

func configJmhtApi() {

	// http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("ok"))
	// })
	// http.HandleFunc("/user/tokeninit", func(w http.ResponseWriter, r *http.Request) {
	// })

	http.HandleFunc("/jmht/api/product/list", CheckLoginStatus(func(w http.ResponseWriter, r *http.Request) {
		var products []g.Product
		err, targetProducts := g.QueryProducts(&products)
		if err != nil {
			log.Println(err)
			RenderDataJson(w, map[string]ResponseStatus{"error": ResponseStatus{Status: 0, Message: err.Error()}})
		} else {
			RenderDataJson(w, map[string][]g.Product{"products": targetProducts})
		}
	}))

	http.HandleFunc("/jmht/api/product/remove", CheckLoginStatus(func(w http.ResponseWriter, r *http.Request) {
		var proR ProductRequest
		ParsePost(&proR, r)

		err := g.DeleteProduct(proR.ID)
		if err != nil {
			log.Println(err)
			RenderDataJson(w, map[string]ResponseStatus{"product": ResponseStatus{Status: 0, Message: "删除失败"}})
		} else {
			RenderDataJson(w, map[string]ResponseStatus{"product": ResponseStatus{Status: 1, Message: "删除成功"}})
		}
	}))

	http.HandleFunc("/jmht/api/product/add", CheckLoginStatus(func(w http.ResponseWriter, r *http.Request) {
		var pR ProductRequest
		ParsePost(&pR, r)
		pro := new(g.Product)
		pro.Name = pR.Name
		pro.Description = pR.Description
		pro.ProductCategoryID = pR.CategoryID

		err := g.CreateProduct(pro)
		if err != nil {
			log.Println(err)
			RenderDataJson(w, map[string]ResponseStatus{"product": ResponseStatus{Status: 0, Message: err.Error()}})
		} else {
			RenderDataJson(w, map[string]g.Product{"product": *pro})
		}
	}))

	http.HandleFunc("/jmht/api/article/list", CheckLoginStatus(func(w http.ResponseWriter, r *http.Request) {
		var articles []g.Article
		if err, articles2 := g.Articles(&articles); err != nil {
			RenderDataJson(w, map[string]ResponseStatus{"list": ResponseStatus{Status: 0, Message: err.Error()}})
		} else {
			RenderDataJson(w, map[string][]g.Article{"list": articles2})
		}
	}))

	http.HandleFunc("/jmht/api/product/category/add", CheckLoginStatus(func(w http.ResponseWriter, r *http.Request) {
		var category CategoryRequest
		ParsePost(&category, r)
		c := new(g.ProductCategory)
		c.Name = category.Name
		err := g.CreateCategory(c)
		if err != nil {
			RenderDataJson(w, map[string]ResponseStatus{"category": ResponseStatus{Status: 0, Message: "添加失败"}})
		} else {
			RenderDataJson(w, map[string]ResponseStatus{"category": ResponseStatus{Status: 1, Message: "添加成功"}})
		}
	}))

	http.HandleFunc("/jmht/api/product/category/list", CheckLoginStatus(func(w http.ResponseWriter, r *http.Request) {
		var categories []g.ProductCategory
		err := g.GetAllCategories(&categories)
		if err != nil {
			RenderDataJson(w, map[string]ResponseStatus{"category": ResponseStatus{Status: 0, Message: "获取失败"}})
		} else {
			RenderDataJson(w, map[string][]g.ProductCategory{"list": categories})
		}
	}))

	http.HandleFunc("/jmht/api/product/category/remove", CheckLoginStatus(func(w http.ResponseWriter, r *http.Request) {
		var categoryR CategoryRequest
		ParsePost(&categoryR, r)
		err := g.DeleteCategory(categoryR.ID)
		if err != nil {
			RenderDataJson(w, map[string]ResponseStatus{"remove": ResponseStatus{Status: 0, Message: "删除失败"}})
		} else {
			RenderDataJson(w, map[string]ResponseStatus{"remove": ResponseStatus{Status: 1, Message: "删除成功"}})
		}
	}))

	http.HandleFunc("/jmht/api/article/add", CheckLoginStatus(func(w http.ResponseWriter, r *http.Request) {

		var articleR ArticleRequest

		ParsePost(&articleR, r)
		// log.Println(articleR)

		article := new(g.Article)
		article.UserID = articleR.UserID
		article.Title = articleR.Title
		article.Content = articleR.Content

		if ok := g.AddArticle(article); !ok {

			RenderDataJson(w, map[string]ResponseStatus{"add": ResponseStatus{Status: 0, Message: "增加失败!"}})
		} else {

			RenderDataJson(w, map[string]g.Article{"article": *article})
		}
	}))

	http.HandleFunc("/jmht/api/uploadImage", CheckLoginStatus(func(w http.ResponseWriter, r *http.Request) {
		err := uploadImg(r)
		if err != nil {
			// log.Fatal(err)
			log.Println(err)
			RenderDataJson(w, map[string]ResponseStatus{"upload": ResponseStatus{Status: 0, Message: "上传失败"}})
		} else {
			RenderDataJson(w, map[string]ResponseStatus{"upload": ResponseStatus{Status: 1, Message: "上传成功"}})
		}
	}))

	http.HandleFunc("/jmht/api/article/update", CheckLoginStatus(func(w http.ResponseWriter, r *http.Request) {
		var articleR ArticleRequest

		ParsePost(&articleR, r)
		log.Println(articleR)

		article := new(g.Article)
		article.ID = articleR.ID
		article.Title = articleR.Title
		article.Content = articleR.Content

		if err := g.UpdateArticle(articleR.ID, article); err != nil {
			RenderDataJson(w, map[string]ResponseStatus{"article": ResponseStatus{Status: 0, Message: "更新失败!"}})
		} else {
			RenderDataJson(w, map[string]ResponseStatus{"article": ResponseStatus{Status: 1, Message: "更新成功!"}})
		}
	}))

	http.HandleFunc("/jmht/api/article/addLike", CheckLoginStatus(func(w http.ResponseWriter, r *http.Request) {
		var articleR ArticleRequest

		ParsePost(&articleR, r)
		log.Println(articleR)

		if err := g.UpdateLikeById(articleR.ID, articleR.Like); err != nil {
			RenderDataJson(w, map[string]ResponseStatus{"add": ResponseStatus{Status: 0, Message: "更新失败!"}})
		} else {
			RenderDataJson(w, map[string]ResponseStatus{"add": ResponseStatus{Status: 1, Message: "更新成功!"}})
		}
	}))

	// { id: 1 }
	http.HandleFunc("/jmht/api/article/delete", CheckLoginStatus(func(w http.ResponseWriter, r *http.Request) {
		var articleR ArticleRequest

		ParsePost(&articleR, r)
		log.Println(articleR)

		if ok := g.DeleteArticle(articleR.ID); ok != nil {
			RenderDataJson(w, map[string]ResponseStatus{"add": ResponseStatus{Status: 0, Message: "删除失败!"}})
		} else {
			RenderDataJson(w, map[string]ResponseStatus{"add": ResponseStatus{Status: 1, Message: "删除成功!"}})
		}
	}))

	http.HandleFunc("/jmht/api/user/logout", func(w http.ResponseWriter, r *http.Request) {

		token := getHeaderToken(r)
		if ok := g.DeleteToken(token); !ok {
			RenderDataJson(w, map[string]ResponseStatus{"logout": ResponseStatus{Status: 0, Message: "登出失败"}})
		} else {
			RenderDataJson(w, map[string]ResponseStatus{"logout": ResponseStatus{Status: 1, Message: "登出成功"}})
		}
	})

	http.HandleFunc("/jmht/api/user/register", func(w http.ResponseWriter, r *http.Request) {
		user := new(g.User)
		user.Name = "jim"
		user.NickName = "张乐"
		user.BirthDay = getTime("1991-10-01")

		pass := new(g.Password)
		pass.PasswordMd5 = "123456"

		user.Password = pass

		// g.DB.Create(&g.User{Name: "jamesduan", NickName: "james"})
		db := g.CreateUser(user)
		if db.RowsAffected == 0 {
			log.Println(db.Error.Error())
			RenderDataJson(w, map[string]string{"message": "创建失败!"})
		} else {
			RenderDataJson(w, map[string]int64{"affected": db.RowsAffected})
		}
		// fmt.Println(db.RowsAffected)
	})

	http.HandleFunc("/jmht/api/user/login", func(w http.ResponseWriter, r *http.Request) {
		// log.Println(r.Cookie("token"))
		var loginedUser g.User
		r.ParseForm()
		// fmt.Println(r.Form)
		body, _ := ioutil.ReadAll(r.Body)
		user := new(User)
		if err := json.Unmarshal(body, user); err != nil {
			// fmt.Println(err)
			RenderDataJson(w, map[string]string{"errorMsg": err.Error()})
		} else {
			if user.UserName != "" && user.Password != "" {
				err := g.CheckLogin(user.UserName, user.Password, &loginedUser)
				if err != nil {
					RenderMsgJson(w, err.Error())
				} else {
					//设置cookie
					//设置cookie，有效期为一年
					// cookiearr := make([]*http.Cookie, 0)
					// cookiearr = append(cookiearr, getCookie("username", loginedUser.Name), getCookie("nickname", loginedUser.NickName), getCookie("token", g.GetToken(loginedUser.ID)), getCookie("user_id", strconv.Itoa(loginedUser.ID)), getCookie("role", loginedUser.RoleName))
					// setCookies(w, cookiearr)
					RenderDataJson(w, map[string]g.User{"user": loginedUser})
				}
			} else {
				RenderMsgJson(w, "数据格式不正确")
			}
		}
		// RenderDataJson(w, loginedUser)
	})

	// http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
	// 	RenderDataJson(w, map[string]string{"Version": g.VERSION})
	// 	// w.Write([]byte(g.VERSION))
	// })

	// http.HandleFunc("/workdir", func(w http.ResponseWriter, r *http.Request) {
	// 	RenderDataJson(w, map[string]string{"Workdir": file.SelfDir()})
	// })

	// http.HandleFunc("/config/reload", func(w http.ResponseWriter, r *http.Request) {
	// 	if strings.HasPrefix(r.RemoteAddr, "127.0.0.1") {
	// 		g.ParseConfig(g.ConfigFile)
	// 		RenderDataJson(w, g.Config())
	// 	} else {
	// 		// w.Write([]byte("no privilege"))
	// 		RenderDataJson(w, map[string]string{"Permission": "no privilege"})
	// 	}
	// })
}
