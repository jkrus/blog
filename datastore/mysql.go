package datastore

import (
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"blog/config"
)

func mySQLConnect(cfg *config.Config) (*gorm.DB, error) {
	strConnect := cfg.DB.User + ":" + cfg.DB.Pass +
		"@tcp([" + cfg.DB.Host + "]:" + strconv.Itoa(cfg.DB.Port) + ")/" + cfg.DB.Name +
		"?charset=" + cfg.DB.Code +
		"&parseTime=True"

	return gorm.Open(mysql.Open(strConnect), &gorm.Config{})
}
