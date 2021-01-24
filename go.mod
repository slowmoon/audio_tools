module github.com/slowmoon/audio_tools

go 1.14

require (
	github.com/RogerLiNing/douyin_watermark_remover v0.0.0-20200918020749-17530283fb84
	github.com/go-kratos/kratos v0.6.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.3
	github.com/google/wire v0.4.0
	github.com/jinzhu/copier v0.1.0
	github.com/slowmoon/base v0.0.0-20201220071941-d1eb07514de0
	golang.org/x/net v0.0.0-20201216054612-986b41b23924 // indirect
	golang.org/x/sys v0.0.0-20201218084310-7d0127a74742 // indirect
	golang.org/x/text v0.3.4 // indirect
	google.golang.org/genproto v0.0.0-20201214200347-8c77b98c765d
	google.golang.org/grpc v1.34.0
	gorm.io/gorm v1.20.8
)

replace google.golang.org/grpc v1.34.0 => google.golang.org/grpc v1.26.0
