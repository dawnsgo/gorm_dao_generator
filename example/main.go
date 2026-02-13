package main

import (
	"context"
	"log"
	"time"

	"github.com/dawnsgo/gorm_dao_generator/example/dao"
	"github.com/dawnsgo/gorm_dao_generator/example/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// PostgreSQL DSN
	dsn := "host=127.0.0.1 port=5432 user=postgres password=123456 dbname=game sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect postgres server failed: %v", err)
	}

	mailDao := dao.NewMail(db)
	baseCtx := context.Background()

	_, err = mailDao.Insert(baseCtx, &model.Mail{
		Title:    "gorm_dao_generator introduction",
		Content:  "The gorm_dao_generator is a tool for automatically generating PostgreSQL Data Access Object.",
		Sender:   1,
		Receiver: 2,
		Status:   1,
		SendTime: time.Now(),
	})
	if err != nil {
		log.Fatalf("failed to insert into postgres database: %v", err)
	}

	mail, err := mailDao.FindOne(baseCtx, func(cols *dao.MailColumns) interface{} {
		return map[string]interface{}{
			cols.Receiver: 2,
		}
	})
	if err != nil {
		log.Fatalf("failed to find a row of data from postgres database: %v", err)
	}

	log.Printf("%+v", mail)
}
