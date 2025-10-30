package health

import (
	"github.com/labstack/echo/v4"
)

type Handler interface {
	CheckCORS(c echo.Context) error
	CheckHealth(c echo.Context) error
}
