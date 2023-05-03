package http

import (
	"fmt"
	infra "github.com/GuilhermeAlegre/CleanArch-Infra"
	v1 "github.com/GuilhermeAlegre/CleanArch-Infra/http/routes/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type http struct {
	router      *fiber.App
	config      Config
	handlers    []infra.IHandler
	middlewares []infra.Middleware
}

func New(config Config) infra.IHttp {
	app := fiber.New()
	app.Use(
		cors.New(),
		logger.New(),
	)

	return &http{
		router: app,
		config: config,
	}
}

func (h *http) Init() error {

	// subscribe middlewares
	for _, middleware := range h.GetMiddlewares() {
		h.router.Use(middleware.Path, middleware.Handler)
	}

	// subscribe routes
	h.subscribeRoutes()

	// open server
	err := h.router.Listen(fmt.Sprintf(":%d", h.config.Port))
	return err
}

func (h *http) GetRouter() *fiber.App {
	return h.router
}

func (h *http) GetHandlers() []infra.IHandler {
	return h.handlers
}

func (h *http) GetMiddlewares() []infra.Middleware {
	return h.middlewares
}

func (h *http) SetHandler(handler infra.IHandler) infra.IHttp {
	h.handlers = append(h.handlers, handler)
	return h
}

func (h *http) SetMiddlewares(middleware ...infra.Middleware) infra.IHttp {
	h.middlewares = append(h.middlewares, middleware...)
	return h
}

func NewMiddleware(path string, handler func(c *fiber.Ctx) error) infra.Middleware {
	return infra.Middleware{
		Path:    path,
		Handler: handler,
	}
}

func (h *http) subscribeRoutes() {
	v1.Routes(h)
}
