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

	//user profile
	profile := e.Group("/profile")
	profile.Use(s.authenticationRequiredMiddleware)
	profile.GET("", s.profileForm)
	profile.POST("", s.profileSubmit)

	//settings
	settings := e.Group("/settings")
	settings.Use(s.authenticationRequiredMiddleware)
	settings.Use(s.settingsRequiredMiddleware)
	settings.GET("", s.settingsGeneralForm)
	settings.POST("", s.settingsGeneralSubmit)
	settings.GET("/:platform", s.settingsByPlatformForm)
	settings.POST("/:platform", s.settingsByPlatformSubmit)

	//users
	users := e.Group("/users")
	users.Use(s.authenticationRequiredMiddleware)
	users.Use(s.userControlAccessRequiredMiddleware)
	users.GET("", s.users)
	users.GET("/new", s.userNewForm)
	users.POST("/new", s.userNewSubmit)

	usersEdit := users.Group("/edit/:user_id")
	usersEdit.Use(s.userControlRequiredMiddleware)
	usersEdit.GET("", s.userEditForm)
	usersEdit.POST("", s.userEditSubmit)

	usersDelete := users.Group("/delete/:user_id")
	usersDelete.Use(s.userControlRequiredMiddleware)
	usersDelete.GET("", s.userDeleteForm)
	usersDelete.POST("", s.userDeleteSubmit)

	return s
}

func (s *Server) Start() error {
	return s.e.Start(s.config.Listen)
}
