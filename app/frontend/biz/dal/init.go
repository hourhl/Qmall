package dal

import (
	"github.com/hourhl/Qmall/app/frontend/biz/dal/mysql"
	"github.com/hourhl/Qmall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
