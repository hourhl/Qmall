package dal

import (
	"github.com/hourhl/Qmall/app/order/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
