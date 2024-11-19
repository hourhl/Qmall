package dal

import (
	"github.com/hourhl/Qmall/app/cart/biz/dal/mysql"
	"github.com/hourhl/Qmall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
