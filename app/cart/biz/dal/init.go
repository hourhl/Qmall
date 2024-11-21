package dal

import (
	"github.com/hourhl/Qmall/app/cart/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
