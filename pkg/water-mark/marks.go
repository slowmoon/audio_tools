package water_mark

import (
    "context"
    "sync"
)

var factory  WaterMarkFactory = WaterMarkFactory{
	marks: make(map[Code]WaterMark),
}


type Code string
const (
	Douyin   Code = "douyin"
)

type WaterMark interface {
    GetVideo(ctx context.Context, origin string) (string, error)
}

type WaterMarkFactory struct {
	lock   sync.Mutex
	marks   map[Code]WaterMark
}

func (w WaterMarkFactory)Get(t Code)WaterMark {
	w.lock.Lock()
	defer w.lock.Unlock()
	return  w.marks[t]
}

func register(name Code, t WaterMark) {
	factory.lock.Lock()
	defer factory.lock.Unlock()
	factory.marks[name]	 = t
}

func GetWaterMark(t Code) WaterMark {
	return  factory.Get(t)
}


