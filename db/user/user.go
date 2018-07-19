package user

import (
	"fmt"
	"jmht-api/common/model"
	"jmht-api/db"
	"log"
)

func Create(user *model.User) error {
	sql := fmt.Sprintf(
		"insert into user(name, email, sex, portrait_url, create_time, update_time) values ('%s', '%s', '%s', '%s', '%s', '%s')",
		user.Name,
		user.Email,
		user.Sex,
		user.PortraitURL,
		user.CreateTime,
		user.UpdateTime,
	)
	// log.Println(sql)
	_, err := db.DB.Exec(sql)
	if err != nil {
		log.Println("exec", sql, "fail", err)
		return err
	}
	return nil
}

func Update(user *model.User) error {
	sql := fmt.Sprintf(
		"update user set name=%s, email=%s, sex=%s, portrait_url=%s, create_time=%s, update_time=%s where id=%d",
		user.Name,
		user.Email,
		user.Sex,
		user.PortraitURL,
		user.CreateTime,
		user.UpdateTime,
		user.ID,
	)
	log.Println(sql)
	result, err := db.DB.Exec(sql)
	log.Println(result.RowsAffected())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Delete(userId int) error {
	sql := fmt.Sprintf("delete from user where id=%d", userId)
	_, err := db.DB.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func SelectUserById(userId int, data *model.User) error {
	sql := fmt.Sprintf("select * from user where id=%d", userId)
	// log.Println(sql)
	rows, err := db.DB.Query(sql)
	log.Println(rows)
	// for row := range rows.Next() {
	// 	log.Println(row)
	// }
	// var userList = make([]*model.User, 0)
	for rows.Next() {
		var (
			id          int
			name        string
			email       string
			sex         string
			portraitURL string
			createTime  string
			updateTime  string
		)
		err := rows.Scan(&id, &name, &email, &sex, &portraitURL, &createTime, &updateTime)
		if err != nil {
			log.Println(err)
		}
		data.ID = id
		data.Name = name
		data.Email = email
		data.Sex = sex
		data.PortraitURL = portraitURL
		data.CreateTime = createTime
		data.UpdateTime = updateTime
		// log.Println(id, name, email, sex, portraitURL, createTime, updateTime)
	}
	if err != nil {
		log.Println("exec", sql, "fail", err)
		return err
	}
	defer rows.Close()
	return nil
}

func selectUserByName(username string) {

}
