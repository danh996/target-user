package main

import (
	"context"
	"log"
	"net"
	"time"

	"user/internal/delivery/grpc"
	"user/internal/domain"
	"user/internal/repository"
	"user/pb"

	"github.com/danh996/go-shop-kit/grpc_server"
	"github.com/danh996/go-shop-kit/token"

	mongoC "user/internal/repository/mongo"

	"github.com/urfave/cli"
	"go.mongodb.org/mongo-driver/mongo"
	mgoOption "go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/reflection"
)

var (
	//MongoCollectionUser collection of user
	MongoCollectionProduct = "users"
)

type srv struct {
	cfg *Config

	mgoDB *mongo.Database

	grpcServer *grpc_server.GRPCServer

	authenticator token.Authenticator

	// load store
	userRepository repository.UserRepository

	// load domain
	userDomain domain.UserDomain
	authDomain domain.AuthDomain
	// define for logger

	authDelivery pb.AuthServiceServer
	userDelivery pb.UserServiceServer

	// parameter
	maxSMSPerDay int
	otpLength    int
	tokenLength  int

	// topic
	postNotificationTopic string
}

func (s *srv) loadConfig(ctx *cli.Context) error {
	s.cfg = &Config{
		HTTP: ConnAddress{
			Host: ctx.GlobalString(HTTPHostFlag.GetName()),
			Port: ctx.GlobalString(HTTPPortFlag.GetName()),
		},
		GRPC: ConnAddress{
			Host: ctx.GlobalString(GRPCHostFlag.GetName()),
			Port: ctx.GlobalString(GRPCPortFlag.GetName()),
		},
		Mongo: Mongo{
			Host:     ctx.GlobalString(MongoHostFlag.GetName()),
			Port:     ctx.GlobalString(MongoPortFlag.GetName()),
			Database: ctx.GlobalString(MongoDatabaseFlag.GetName()),
			Username: ctx.GlobalString(MongoDatabaseUsernameFlag.GetName()),
			Password: ctx.GlobalString(MongoDatabasePasswordFlag.GetName()),
		},
		TokenKey: ctx.GlobalString(TokenKeyFlag.GetName()),
	}

	return nil
}

func (s *srv) connectMongo() error {
	mgoClientOptions := mgoOption.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	var err error
	mgoClient, err := mongo.Connect(context.TODO(), mgoClientOptions)
	if err != nil {
		return err
	}
	s.mgoDB = mgoClient.Database(s.cfg.Mongo.Database)
	log.Println("connect mongodb success")
	return nil
}

func (s *srv) loadRepository() error {
	s.userRepository = mongoC.NewUserRepository(s.mgoDB.Collection(MongoCollectionProduct))

	log.Println("load repository success")
	return nil
}

func (s *srv) loadDomain() error {
	s.userDomain = domain.NewUserDomain(s.userRepository)
	s.authDomain = domain.NewAuthDomain(
		s.userRepository,
		s.authenticator,
	)
	log.Println("load domain success")
	log.Println(s.authDomain)
	return nil
}

func (s *srv) loadDelivery() error {

	s.userDelivery = grpc.NewUserDelivery(s.userDomain)
	s.authDelivery = grpc.NewAuthDelivery(s.authDomain, s.authenticator)
	log.Println("load delivery success")
	return nil
}

func (s *srv) loadGRPCServer() error {
	s.grpcServer = grpc_server.NewGRPCServer(s.cfg.ServiceName, s.cfg.GRPC.Host, s.cfg.GRPC.Port)
	s.grpcServer.InitServer()
	pb.RegisterUserServiceServer(s.grpcServer.Server, s.userDelivery)
	pb.RegisterAuthServiceServer(s.grpcServer.Server, s.authDelivery)
	log.Println("load GRPC SERVER success")
	return nil
}

func (s *srv) startGRPCServer() {
	ln, err := net.Listen("tcp", s.cfg.GRPC.String())
	if err != nil {
		panic(err)
	}
	log.Printf("START GRPC SERVER success on port %v", s.cfg.GRPC.String())

	reflection.Register(s.grpcServer.Server)
	if err := s.grpcServer.Server.Serve(ln); err != nil {
		panic(err)
	}

}

func (s *srv) loadAuthenticator() error {

	log.Println(s.cfg.TokenKey)
	log.Println("loading authentication")
	var err error
	s.authenticator, err = token.NewJWTAuthenticator(s.cfg.TokenKey, 15*24*time.Hour)
	return err
}
