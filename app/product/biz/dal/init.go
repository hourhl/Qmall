package dal

import (
	"github.com/hourhl/Qmall/app/product/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
