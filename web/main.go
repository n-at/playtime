package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

var (
	AssetsWebRoot = "/assets"
	AssetsRoot    = "assets"
)

func New() *echo.Echo {
	e := echo.New()
	e.Renderer = pongo2Renderer{}
	e.HTTPErrorHandler = httpErrorHandler
	e.Static(AssetsWebRoot, AssetsRoot)
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(log.Fields{
				"method": values.Method,
				"uri":    values.URI,
				"status": values.Status,
			}).Info("request")
			return nil
		},
	}))

	//routes

	//

	return e
}
