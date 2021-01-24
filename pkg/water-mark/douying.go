package water_mark

import (
	"context"
	douyin "github.com/RogerLiNing/douyin_watermark_remover"
	"strings"
)

type douying struct {

}

func (d douying) GetVideo(ctx context.Context, origin string) (string, error) {
	link, err :=  douyin.WatermarkRemover(origin)
	if err != nil {
		return "", err
	}
	return  strings.Replace(link, "http://", "https://", 1), nil
}

func init()  {
	register(Douyin, douying{})
}

