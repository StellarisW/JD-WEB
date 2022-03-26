package boot

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main/app/global"
	"time"
)

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

func MySQLSetup() {
	config := g.Config.Mysql
	db, err := gorm.Open(mysql.Open(config.Dsn()))
	if err != nil {
		g.Logger.Fatalf("Initialize MySQL server failed, err: %v\n", err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxIdleTime(10 * time.Second)  // 最大空闲时间
	sqlDB.SetConnMaxLifetime(100 * time.Second) // 最大存活时间
	sqlDB.SetMaxIdleConns(10)                   // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)                  // 最大连接数
	err = sqlDB.Ping()
	if err != nil {
		g.Logger.Fatalf("Connect to MySQL server failed, err: %v\n", err)
	}
	g.Logger.Info("Initialize MySQL server successfully!")
	g.DB = db
}
