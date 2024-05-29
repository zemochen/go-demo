package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/zemochen/go-demo/gomall/demo/demo_proto/biz/dal"
	"github.com/zemochen/go-demo/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/zemochen/go-demo/gomall/demo/demo_proto/biz/model"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()

	// result := mysql.DB.Create(&model.User{Email: "demo@gomall.com", Password: "pass"}) // pass pointer of data to Create

	// 	user.ID             // returns inserted data's primary key
	// 	result.Error        // returns error
	// 	result.RowsAffected // returns inserted records count

	// var users = []User{{Email: "zemo_1@gomall.com", Password: "pass_1"}, ...., {Email: "zemo_1000@gomall.com", Password: "pass_1000"}}

	// // batch size 100
	// db.CreateInBatches(users, 100)

	mysql.DB.Model(&model.User{}).Where("email = ?", "demo@gomall.com").Update("Password", "hello")
	var queryUser model.User

	mysql.DB.Model(&model.User{}).Where("email = ?", "demo@gomall.com").First(&queryUser)

	fmt.Printf("user: %+v\n", queryUser)

	// mysql.DB.Where("email = ?", "demo@gomall.com").Delete(&model.User{})
}
