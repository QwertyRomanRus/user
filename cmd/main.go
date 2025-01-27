package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"user/internal/api/v1/user"
	"user/internal/config"
	"user/pkg/user_v1"
)

const (
	EnvPath = ".env"
)

func main() {
	err := config.Load(EnvPath)
	ctx := context.Background()
	if err != nil {
		log.Fatal("Error loading .env: ", err)
	}
	initPg(&ctx)

	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Error grpc listening on port: ", err)
	}
	defer l.Close()

	userApi := user.Implementation{}
	s := grpc.NewServer()
	user_v1.RegisterUserServiceV1Server(s, &userApi)

	reflection.Register(s)

	err = s.Serve(l)
	if err != nil {
		log.Fatal("Error grpc listening on port: ", err)
	}

	//initGrpc(&ctx)
}

func initPg(ctx *context.Context) {
	pgConf, err := config.NewPGConfig()
	if err != nil {
		log.Fatal("Error loading pg config: ", err)
	}

	pgCon, err := pgxpool.Connect(*ctx, pgConf.GetDsn())
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer pgCon.Close() // todo мб неправильно будет закрываться, проверить

	err = pgCon.Ping(*ctx)
	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}

}

//func initGrpc(ctx *context.Context) {
//
//}
