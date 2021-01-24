package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	pb "github.com/slowmoon/audio_tools/api"
	"github.com/slowmoon/audio_tools/internal/dao"
	water_mark "github.com/slowmoon/audio_tools/pkg/water-mark"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.DemoServer), new(*Service)))

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
}

func (s *Service) GetAudioList(ctx context.Context, e *empty.Empty) (*pb.GetAudioListRes, error) {
	lists, err := s.dao.GetAudioList(ctx)
	if err != nil {
		log.Info("get audio list fail %+v", err)
		return &pb.GetAudioListRes{}, nil
	}
	var audios []*pb.AudioInfo
	copier.Copy(&audios, lists)
	return  &pb.GetAudioListRes{Lists: audios}, nil
}

func (s *Service) GetAudioInfo(ctx context.Context, req *pb.GetAudioInfoReq) (*pb.GetAudioInfoRes, error) {
	//获取  type
	var (
		encoded string
		err error
	)
	switch req.Type {
	case "douyin":
		encoded, err = water_mark.GetWaterMark(water_mark.Douyin).GetVideo(ctx, req.OriginUrl)
	default:
		return  nil, fmt.Errorf("unsupport audio type")
	}
	if len(encoded) == 0 {
		err = fmt.Errorf("get water mark fail")
	}
	return &pb.GetAudioInfoRes{VideoUrl: encoded}, err
}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		ac:  &paladin.TOML{},
		dao: d,
	}
	cf = s.Close
	err = paladin.Watch("application.toml", s.ac)
	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
}
