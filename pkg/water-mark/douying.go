package water_mark

import (
	"context"
	douyin "github.com/RogerLiNing/douyin_watermark_remover"
)

type douying struct {

}

func (d douying) GetVideo(ctx context.Context, origin string) (string, error) {
	return douyin.WatermarkRemover(origin)
}

func init()  {
	register(Douyin, douying{})
}

