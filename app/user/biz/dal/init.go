package dal

import (
	"github.com/hourhl/Qmall/app/user/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
