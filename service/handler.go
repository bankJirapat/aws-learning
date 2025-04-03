package service

import "github.com/labstack/echo/v4"

type ServiceHandler struct {
	svc Service
}

func RegisterRouter(e *echo.Group, svc Service) {
	// h := ServiceHandler{
	// 	svc: svc,
	// }

}
