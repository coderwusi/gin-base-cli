package model

import (
	"fmt"
	"gin-base-cli/setting"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// DB 数据库链接单例

var DB *gorm.DB

func MysqlInit(config *setting.MySQLConfig) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", config.Username, config.Password, config.Host, config.Port, config.Dbname, config.Suffix)
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		zap.L().Fatal("mysql数据库连接错误:", zap.Error(err))
		return
	}

	if err := DB.AutoMigrate(); err != nil {
		zap.L().Fatal("mysql数据库迁移错误:", zap.Error(err))
		return
	}

	fmt.Println("Mysql连接成功")
}
