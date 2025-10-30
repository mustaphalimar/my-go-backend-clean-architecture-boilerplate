package studyplan

import "github.com/labstack/echo/v4"

type Handler interface {
	Get() echo.HandlerFunc
	GetAll() echo.HandlerFunc
}
