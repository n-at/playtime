package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"playtime/storage"
)

type Configuration struct {
	AssetsWebRoot      string
	AssetsRoot         string
	UploadsWebRoot     string
	UploadsRoot        string
	Listen             string
	TemplatesDebug     bool
	TemplatesRoot      string
	TemplatesExtension string
}

type Web struct {
	e       *echo.Echo
	config  *Configuration
	storage *storage.Storage
}

func New(config *Configuration, storage *storage.Storage) *Web {
	e := echo.New()
	e.Renderer = newPongo2Renderer(config)
	e.HTTPErrorHandler = httpErrorHandler
	e.Static(config.AssetsWebRoot, config.AssetsRoot)
	e.Static(config.UploadsWebRoot, config.UploadsRoot)
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

	return &Web{
		e:       e,
		config:  config,
		storage: storage,
	}
}

func (w *Web) Start() error {
	return w.e.Start(w.config.Listen)
}
