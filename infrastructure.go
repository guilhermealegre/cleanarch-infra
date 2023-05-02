package infra

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

type Infra struct {
	validator IValidator
	logger    ILogger
	postgres  IPostgres
	redis     IRedis
	http      IHttp
	services  []IService
}

type Middleware struct {
	Path    string
	Handler func(c *fiber.Ctx) error
}

func New() *Infra {
	return &Infra{}
}

func (i *Infra) Start() error {

	for _, service := range i.services {
		if err := service.Init(); err != nil {
			return err
		}
	}

	return nil
}

func (i *Infra) Stop() {
	os.Exit(1)
}

func (i *Infra) Logger() ILogger {
	return i.logger
}

func (i *Infra) Postgres() IPostgres {
	return i.postgres
}

func (i *Infra) Redis() IRedis {
	return i.redis
}

func (i *Infra) Http() IHttp {
	return i.http
}

func (i *Infra) Validator() IValidator {
	return i.validator
}

func (i *Infra) SetLogger(logger ILogger) IInfra {
	i.services = append(i.services, logger)
	i.logger = logger
	return i
}

func (i *Infra) SetValidator(validator IValidator) IInfra {
	i.services = append(i.services, validator)
	i.validator = validator
	return i
}

func (i *Infra) SetPostgres(postgres IPostgres) IInfra {
	i.services = append(i.services, postgres)
	i.postgres = postgres
	return i
}

func (i *Infra) SetRedis(redis IRedis) IInfra {
	i.services = append(i.services, redis)
	i.redis = redis
	return i
}

func (i *Infra) SetHttp(http IHttp) IInfra {
	i.services = append(i.services, http)
	i.http = http
	return i
}
