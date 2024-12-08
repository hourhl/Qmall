package mysql

import (
	"fmt"
	"github.com/hourhl/Qmall/app/cart/biz/model"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	// dev
	//dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	// unit test
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3304)/cart?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	fmt.Printf("cart dsn %s\n", dsn)
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)

	DB.AutoMigrate(&model.Cart{})
	if err != nil {
		panic(err)
	}
}
