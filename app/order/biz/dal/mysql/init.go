package mysql

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hourhl/Qmall/app/order/biz/model"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	// dev
	//dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	// unit test
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3307)/order?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}
	if err := DB.AutoMigrate(&model.Order{}, &model.OrderItem{}); err != nil {
		fmt.Printf("auto migrate err:%v\n", err)
		klog.Error(err)
	}
}
