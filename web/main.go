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

	e.HEAD(UploadsWebRoot+"*", s.uploadsHead)

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

	//games
	games := e.Group("/games")
	games.Use(s.authenticationRequiredMiddleware)
	games.GET("", s.games)
	games.POST("/upload", s.gameUpload)

	gamesEmulationSettings := games.Group("/emulation-settings/:game_id")
	gamesEmulationSettings.Use(s.gameRequiredMiddleware)
	gamesEmulationSettings.GET("", s.gameEmulationSettingsForm)
	gamesEmulationSettings.POST("", s.gameEmulationSettingsSubmit)

	gamesEdit := games.Group("/edit/:game_id")
	gamesEdit.Use(s.gameRequiredMiddleware)
	gamesEdit.GET("", s.gameEditForm)
	gamesEdit.POST("", s.gameEditSubmit)

	gamesDelete := games.Group("/delete/:game_id")
	gamesDelete.Use(s.gameRequiredMiddleware)
	gamesDelete.GET("", s.gameDeleteForm)
	gamesDelete.POST("", s.gameDeleteSubmit)

	uploadBatch := games.Group("/upload-batch/:upload_batch_id")
	uploadBatch.Use(s.uploadBatchRequiredMiddleware)
	uploadBatch.GET("", s.gameUploadBatchForm)
	uploadBatch.POST("", s.gameUploadBatchSubmit)

	saveStates := games.Group("/save-states/:game_id")
	saveStates.Use(s.gameRequiredMiddleware)
	saveStates.GET("", s.saveStates)

	saveStateDelete := saveStates.Group("/delete/:save_state_id")
	saveStateDelete.Use(s.saveStateRequiredMiddleware)
	saveStateDelete.GET("", s.saveStateDeleteForm)
	saveStateDelete.POST("", s.saveStateDeleteSubmit)

	//play

	play := e.Group("/play/:game_id")
	play.Use(s.authenticationRequiredMiddleware)
	play.Use(s.gameRequiredMiddleware)
	play.GET("", s.play)

	return s
}

func (s *Server) Start() error {
	return s.e.Start(s.config.Listen)
}
