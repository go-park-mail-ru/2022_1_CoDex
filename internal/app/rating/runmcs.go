package mcsrtng

import (
	"codex/internal/pkg/database"
	proto "codex/internal/pkg/rating/delivery/grpc"
	ratrepository "codex/internal/pkg/rating/repository"
	ratusecase "codex/internal/pkg/rating/usecase"
	"codex/internal/pkg/utils/config"
	"codex/internal/pkg/utils/log"

	"net"

	"google.golang.org/grpc"
)

func RunServer() {
	db := database.InitDatabase()
	db.Connect()
	defer db.Disconnect()

	autRep := ratrepository.InitRatRep(db)
	autUsc := ratusecase.InitRatUsc(autRep)

	s := grpc.NewServer()

	proto.RegisterPosterServer(s, autUsc)

	l, err := net.Listen(config.DevConfigStore.Mcs.Rating.ConnType, ":"+config.DevConfigStore.Mcs.Rating.Port)
	if err != nil {
		log.Warn("{RunServer} mcs rtng")
		log.Error(err)
	}

	err = s.Serve(l)
	if err != nil{
		log.Error(err)
	}
	
}
