package mysql

import (
	"fmt"

	"github.com/austoin/GolangStore/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func BuildDSN(conf config.MySQL) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	)
}

func Open(conf config.MySQL) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(BuildDSN(conf)), &gorm.Config{})
}
