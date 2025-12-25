package dao

import (
	"github.com/dawnsgo/gorm_dao_generator/example/dao/internal"
	"gorm.io/gorm"
)

type (
	MailColumns = internal.MailColumns
	MailOrderBy = internal.MailOrderBy
)

type Mail struct {
	*internal.Mail
}

func NewMail(db *gorm.DB) *Mail {
	return &Mail{Mail: internal.NewMail(db)}
}
