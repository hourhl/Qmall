package dal

import (
	"github.com/hourhl/Qmall/app/user/biz/dal/mysql"
	"github.com/hourhl/Qmall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
