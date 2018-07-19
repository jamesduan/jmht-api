package rpc

import (
	"database/sql"
	"fmt"
	"jmht-api/common/model"
	"log"
	"os"
	"testing"
	"time"
)

const RPCServer = "127.0.0.1:6030"

// const Port = 6030

func TestUpdateUser(t *testing.T) {
	timeout := time.Duration(1) * time.Second
	client, err := JsonRpcClient("tcp", RPCServer, timeout)
	defer client.Close()
	if err != nil {
		fmt.Println("dial failed:", err)
		os.Exit(1)
	}
	user := new(model.User)
	user.ID = 2
	user.Name = "'xxxx'"
	user.Sex = "'女'"
	user.PortraitURL = "'////////xx'"
	user.CreateTime = "'2018-03-23 01:21:00'"
	user.UpdateTime = "'2018-03-23 01:21:00'"
	user.Email = "'jamesduanling@gmail.com'"
	result := new(sql.Result)
	client.Call("User.UpdateUser", user, result)
	// log.Println(response)
}

func TestDeleteUser(t *testing.T) {
	timeout := time.Duration(1) * time.Second
	client, err := JsonRpcClient("tcp", RPCServer, timeout)
	defer client.Close()
	if err != nil {
		fmt.Println("dial failed:", err)
		os.Exit(1)
	}
	u := new(model.User)
	client.Call("User.GetUser", 6, &u)
	result := new(sql.Result)
	client.Call("User.DeleteUser", u.ID, result)
}

func TestGetUser(t *testing.T) {
	timeout := time.Duration(1) * time.Second
	client, err := JsonRpcClient("tcp", RPCServer, timeout)
	defer client.Close()
	if err != nil {
		fmt.Println("dial failed:", err)
		os.Exit(1)
	}
	response := new(model.User)
	client.Call("User.GetUser", 5, &response)
	log.Println(response)

}

func TestUser(t *testing.T) {

	timeout := time.Duration(1) * time.Second
	client, err := JsonRpcClient("tcp", RPCServer, timeout)
	defer client.Close()
	if err != nil {
		fmt.Println("dial failed:", err)
		os.Exit(1)
	}
	// db.Init()
	var user = new(model.User)
	user.Name = "张乐"
	user.Email = "jamesduanling@hotmail.com"
	user.PortraitURL = "github.com/jamesduan/123"
	user.Sex = "男"
	user.CreateTime = "2018-3-19 13:56"
	user.UpdateTime = "2018-3-19 13:56"
	response := new(model.User)
	client.Call("User.CreateUser", user, &response)
	// Create(user)
	// user := model.User{Name: "jamesduan", Email: "jamesduanling@hotmail.com", Sex: "男", CreateTime: "2018-3-19 13:56", UpdateTime: "2018-3-19 13:56"}
	// _, err := Create(user)
	// if err != nil {
	// 	log.Println(err)
	// }
}
