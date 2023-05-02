package infra

import (
	"github.com/gofiber/fiber/v2"
)

type Context struct {
	*fiber.Ctx
}

type Session struct {
	UserId    int    `json:"user_id"`
	TokenUUID string `json:"token_uuid"`
}

func GetContext(ctx *fiber.Ctx) *Context {

	return &Context{
		ctx,
	}
}

func (c *Context) SetSession() {
	c.Context().SetUserValue("test", "ola")
}

func (c *Context) GetSession() string {
	return c.Context().UserValue("test").(string)
}
