package dao

import (
	"context"
	"github.com/slowmoon/audio_tools/internal/model"
)

func (d *dao) GetAudioList(ctx context.Context) (lists []model.Audios,err error) {
	err = d.db.Find(&lists).Error
	return
}
