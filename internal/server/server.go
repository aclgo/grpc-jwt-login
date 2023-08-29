package server

import (
	"fmt"
	"net"

	"github.com/aclgo/grpc-jwt/config"
	"github.com/aclgo/grpc-jwt/internal/interceptor"
	sessionUC "github.com/aclgo/grpc-jwt/internal/session/usecase"
	"github.com/aclgo/grpc-jwt/internal/user/delivery/grpc/service"
	userRepo "github.com/aclgo/grpc-jwt/internal/user/repository"
	userUC "github.com/aclgo/grpc-jwt/internal/user/usecase"
	"github.com/aclgo/grpc-jwt/pkg/logger"
	"github.com/aclgo/grpc-jwt/proto"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

type Server struct {
	db          *sqlx.DB
	redisClient *redis.Client
	logger      logger.Logger
	config      *config.Config
}

func NewServer(db *sqlx.DB, redisClient *redis.Client,
	logger logger.Logger, config *config.Config) *Server {
	return &Server{
		db:          db,
		redisClient: redisClient,
		logger:      logger,
		config:      config,
	}
}

func (s *Server) Run() error {
	interceptor := interceptor.NewInterceptor(s.logger)

	sessUC := sessionUC.NewSessionUC(s.logger, s.redisClient, s.config.SecretKey)
	usRepo := userRepo.NewPostgresRepo(s.db)
	usUC := userUC.NewUserUC(s.logger, usRepo, nil, sessUC)

	userService := service.NewUserService(s.logger, usUC)

	listen, err := net.Listen("tcp", "localhost:"+s.config.ServerPort)
	// fmt.Println(s.config.ServerPort)
	if err != nil {
		s.logger.Errorf("net.Listen: %v", err)
	}

	opts := []grpc.ServerOption{
		// grpc.KeepaliveParams(grpc.KeepaliveParams{
		// 	grpc.MaxConnectionIdle:
		// 	TIMeout:
		// MaxConnectionAge:
		// Time:
		// }),
		grpc.UnaryInterceptor(interceptor.Logger),
		// grpc.ChainUnaryInterceptor(
		// 	grpc_ctxtags.UnaryServerInterceptor(),
		// 	grpc_prometheus.UnaryServerInterceptor(),
		// 	grpc_recovery.UnaryServerInterceptor(),
		// ),
	}

	server := grpc.NewServer(opts...)
	proto.RegisterUserServiceServer(server, userService)
	s.logger.Infof("server starting port %s", s.config.ServerPort)
	if err := server.Serve(listen); err != nil {
		return fmt.Errorf("Run.NewServer: %v", err)
	}

	return nil
}
