package dal

import (
	"github.com/hourhl/Qmall/app/payment/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
