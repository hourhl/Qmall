package dal

import (
	"github.com/hourhl/Qmall/app/product/biz/dal/mysql"
	"github.com/hourhl/Qmall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
