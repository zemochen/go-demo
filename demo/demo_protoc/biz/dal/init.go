package dal

import (
	"github.com/zemochen/go-demo/gomall/demo/demo_proto/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
