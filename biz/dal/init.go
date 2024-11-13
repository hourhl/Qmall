package dal

import (
	"github.com/hourhl/Qmall/biz/dal/mysql"
	"github.com/hourhl/Qmall/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
