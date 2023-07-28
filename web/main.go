package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"playtime/storage"
)

const (
	SessionCookieName = "playtime-sess-id"
	AssetsWebRoot     = "/assets"
	UploadsWebRoot    = "/uploads"
)

type Configuration struct {
	AssetsRoot         string
	UploadsRoot        string
	Listen             string
	TemplatesDebug     bool
	TemplatesRoot      string
	TemplatesExtension string
}

type Server struct {
	e       *echo.Echo
	config  *Configuration
	storage *storage.Storage
}

func New(config *Configuration, storage *storage.Storage) *Server {
	e := echo.New()
	e.Renderer = newPongo2Renderer(config)
	e.HTTPErrorHandler = httpErrorHandler
	e.Static(AssetsWebRoot, config.AssetsRoot)
	e.Static(UploadsWebRoot, config.UploadsRoot)
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

	s := &Server{
		e:       e,
		config:  config,
		storage: storage,
	}

	e.Use(s.contextCustomizationMiddleware)

	e.GET("/", s.index)

	//authentication
	e.GET("/login", s.loginForm)
	e.POST("/login", s.loginSubmit)
	e.GET("/logout", s.logout)

	return s
}

func (s *Server) Start() error {
	return s.e.Start(s.config.Listen)
}
