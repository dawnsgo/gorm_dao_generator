package main

import (
	"context"
	"log"
	"time"

	"github.com/dawnsgo/gorm_dao_generator/example/dao"
	"github.com/dawnsgo/gorm_dao_generator/example/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// MySQL 8.0+ DSN with allowNativePasswords for authentication compatibility
	dsn := "root:123456@tcp(127.0.0.1:3306)/game?charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true"

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}))
	if err != nil {
		log.Fatalf("connect mysql server failed: %v", err)
	}

	mailDao := dao.NewMail(db)
	baseCtx := context.Background()

	_, err = mailDao.Insert(baseCtx, &model.Mail{
		Title:    "gorm_dao_generator introduction",
		Content:  "The gorm_dao_generator is a tool for automatically generating Mysql Data Access Object.",
		Sender:   1,
		Receiver: 2,
		Status:   1,
		SendTime: time.Now(),
	})
	if err != nil {
		log.Fatalf("failed to insert into mysql database: %v", err)
	}

	mail, err := mailDao.FindOne(baseCtx, func(cols *dao.MailColumns) interface{} {
		return map[string]interface{}{
			cols.Receiver: 2,
		}
	})
	if err != nil {
		log.Fatalf("failed to find a row of data from mysql database: %v", err)
	}

	log.Printf("%+v", mail)
}
