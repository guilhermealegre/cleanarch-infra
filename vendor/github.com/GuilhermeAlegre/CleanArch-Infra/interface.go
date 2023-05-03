package infra

import (
	"context"
	redis_cli "github.com/go-redis/redis/v7"
	"github.com/gocraft/dbr/v2"
	"github.com/gofiber/fiber/v2"
)

type IInfra interface {
	Start() error
	Stop()
	Postgres() IPostgres
	Validator() IValidator
	Redis() IRedis
	Logger() ILogger
	Http() IHttp
	SetPostgres(postgres IPostgres) IInfra
	SetValidator(validator IValidator) IInfra
	SetLogger(logger ILogger) IInfra
	SetRedis(redis IRedis) IInfra
	SetHttp(http IHttp) IInfra
}

type IHandler interface {
	Get(ctx *fiber.Ctx) error
	Put(ctx *fiber.Ctx) error
	Post(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type IRedis interface {
	IService
	Conn() *redis_cli.Client
}

type IPostgres interface {
	IService
	Read(ctx context.Context) *dbr.Session
	Write(ctx context.Context) *dbr.Session
}

type IValidator interface {
	IService
}

type ILogger interface {
	IService
}

type IHttp interface {
	IService
	GetRouter() *fiber.App
	GetHandlers() []IHandler
	GetMiddlewares() []Middleware
	SetHandler(handler IHandler) IHttp
	SetMiddlewares(middleware ...Middleware) IHttp
}

type IService interface {
	Init() error
}
