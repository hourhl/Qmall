package dal

import (
	"github.com/hourhl/Qmall/app/payment/biz/dal/mysql"
	"github.com/hourhl/Qmall/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
