package dao

import (
	"github.com/slowmoon/audio_tools/internal/model"
	"github.com/slowmoon/base/databases"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
)

func NewDB() (db *gorm.DB, cf func(), err error) {
	var (
		cfg databases.DbConfig
		ct paladin.TOML
	)
	if err = paladin.Get("db.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
		return
	}
	db , err= databases.NewMySqlDb(cfg)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(model.Audios{})
	cf = func() {}
	return
}

