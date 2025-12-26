package handlers

import (
	"api-gateway/handlers/notification"
	"api-gateway/handlers/order"
	"api-gateway/handlers/payment"
	"api-gateway/handlers/product"
	"api-gateway/handlers/user"

	"github.com/labstack/echo/v4"
)

func RegisterPublicRoutes(g *echo.Group) {
	user.RegisterProtectedRoutes(g)
	product.RegisterProtectedRoutes(g)
}

func RegisterProtectedRoutes(g *echo.Group) {
	user.RegisterPublicRoutes(g)
	product.RegisterPublicRoutes(g)

	order.RegisterRoutes(g)
	payment.RegisterRoutes(g)

	notification.RegisterRoutes(g)
}

func RegisterAllRoutes(g *echo.Group, jwtMiddleware echo.MiddlewareFunc) {
	RegisterPublicRoutes(g)

	protected := g.Group("")
	protected.Use(jwtMiddleware)
	RegisterProtectedRoutes(protected)
}
