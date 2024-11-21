package dal

import (
	"github.com/hourhl/Qmall/app/checkout/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
