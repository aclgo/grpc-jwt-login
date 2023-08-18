package server

import (
	"github.com/aclgo/grpc-jwt/config"
	"github.com/aclgo/grpc-jwt/pkg/logger"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
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
	return nil
}
