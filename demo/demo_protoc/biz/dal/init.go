package dal

import (
	"github.com/zemochen/go-demo/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/zemochen/go-demo/gomall/demo/demo_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
