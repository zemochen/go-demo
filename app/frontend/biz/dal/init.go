package dal

import (
	"github.com/zemochen/go-demo/gomall/app/frontend/biz/dal/mysql"
	"github.com/zemochen/go-demo/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
