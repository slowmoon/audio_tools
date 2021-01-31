package water_mark

import (
	"context"
	douyin "github.com/RogerLiNing/douyin_watermark_remover"
)

type douying struct {

}

func (d douying) GetVideo(ctx context.Context, origin string) (string, error) {
	link, err :=  douyin.WatermarkRemover(origin)
	if err != nil {
		return "", err
	}
	return  link, nil
}



func init()  {
	register(Douyin, douying{})
}

