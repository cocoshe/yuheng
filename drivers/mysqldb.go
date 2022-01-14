package drivers

import (
	"backend/config"
	"backend/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var MysqlDb *gorm.DB

func init() {
	// get db config
	dbDSN := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GlobalConfig.GetString("db.User"),
		config.GlobalConfig.GetString("db.Pass"),
		config.GlobalConfig.GetString("db.Addr"),
		config.GlobalConfig.GetString("db.Port"),
		config.GlobalConfig.GetString("db.DB"),
	)

	// connect and open db connection
	var err error
	MysqlDb, err = gorm.Open("mysql", dbDSN)
	if err != nil {
		log.Fatal("Mysql connect fail!")
	}
	MysqlDb.SingularTable(true)

	// max open connections
	dbMaxOpenConns := config.GlobalConfig.GetInt("DB_MAX_OPEN_CONNS")
	MysqlDb.DB().SetMaxOpenConns(dbMaxOpenConns)

	// max idle connections
	dbMaxIdleConns := config.GlobalConfig.GetInt("DB_MAX_IDLE_CONNS")
	MysqlDb.DB().SetMaxIdleConns(dbMaxIdleConns)

	// max lifetime of connection if <=0 will forever
	dbMaxLifetimeConns := config.GlobalConfig.GetInt("DB_MAX_LIFETIME_CONNS")
	MysqlDb.DB().SetConnMaxLifetime(time.Duration(dbMaxLifetimeConns))

	// check db connection at once avoid connect failed
	// else error will be reported until db first sql operate
	if MysqlDbErr := MysqlDb.DB().Ping(); nil != MysqlDbErr {
		panic("database connect failed: " + MysqlDbErr.Error())
	}

	MysqlDb.AutoMigrate(
		&models.User{},
		&models.Accusation{},
		&models.Appeal{},
	)

}
