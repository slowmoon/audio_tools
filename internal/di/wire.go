// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/slowmoon/audio_tools/internal/dao"
	"github.com/slowmoon/audio_tools/internal/server/grpc"
	"github.com/slowmoon/audio_tools/internal/server/http"
	"github.com/slowmoon/audio_tools/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
