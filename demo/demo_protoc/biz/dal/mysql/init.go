package mysql

import (
	"fmt"
	"os"

	"github.com/zemochen/go-demo/gomall/demo/demo_proto/biz/model"
	"github.com/zemochen/go-demo/gomall/demo/demo_proto/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB"),
	)
	// fmt.Printf("dsn connection:%s", dsn)

	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	type Version struct {
		Version string
	}

	var v Version

	err = DB.Raw("select version() as version").Scan(&v).Error
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&model.User{})
	fmt.Println("MYSQL Version:", v)
	// fmt.Printf("%#v", DB.Debug().Exec("select version()"))
}
