package v1

import (
	"auth/handler/http/v1/authentication"
	"auth/infra"
)

func Routes(httpService infra.IHttp) {
	v1 := httpService.GetRouter().Group("/v1")
	for _, h := range httpService.GetHandlers() {

		switch handler := h.(type) {
		case *authentication.AuthenticationHandler:
			v1.Post("/auth/get-token", handler.GetToken)
			v1.Post("/auth/refresh-token", handler.RefreshToken)
			v1.Post("/auth/sign-up", handler.SignUp)
			v1.Post("/auth/log-out", handler.LogOut)
		}
	}
}
