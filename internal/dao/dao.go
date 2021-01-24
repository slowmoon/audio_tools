package dao

import (
	"context"
	"github.com/google/wire"
	"github.com/slowmoon/audio_tools/internal/model"
	"gorm.io/gorm"
)

var Provider = wire.NewSet(New, NewDB)

//go:generate kratos tool genbts
// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	GetAudioList(ctx context.Context) (lists []model.Audios,err error)
}

// dao dao.
type dao struct {
	db          *gorm.DB
}



// New new a dao and return.
func New( db *gorm.DB) (d Dao, cf func(), err error) {
	return  &dao{db: db}, func() {}, nil
}

//func newDao( db *gorm.DB) (d *dao, cf func(), err error) {
//	var cfg struct{
//		DemoExpire xtime.Duration
//	}
//	if err = paladin.Get("application.toml").UnmarshalTOML(&cfg); err != nil {
//		return
//	}
//	d = &dao{
//		db: db,
//		redis: r,
//		mc: mc,
//		cache: fanout.New("cache"),
//		demoExpire: int32(time.Duration(cfg.DemoExpire) / time.Second),
//	}
//	cf = d.Close
//	return
//}

// Close close the resource.
func (d *dao) Close() {
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}
