package dal

import "github.com/hourhl/Qmall/app/auth/biz/dal/redis"

func Init() {
	redis.Init()
	//mysql.Init()
}
