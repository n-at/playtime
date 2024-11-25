package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"playtime/storage"
	"playtime/web/gamesession"
	"playtime/web/heartbeatpool"
	"sync"
	"time"
)

const (
	SessionCookieName = "playtime-sess-id"
	AssetsWebRoot     = "/assets"
	UploadsWebRoot    = "/uploads"
	HeartbeatInterval = 10 * time.Second
	HeartbeatThreads  = 4
)

type Configuration struct {
	AssetsRoot  string
	UploadsRoot string
	Listen      string

	//templates
	TemplatesDebug     bool
	TemplatesRoot      string
	TemplatesExtension string

	//emulator
	EmulatorDebug bool

	//netplay
	NetplayEnabled     bool
	NetplayDebug       bool
	TurnServerUrl      string
	TurnServerUser     string
	TurnServerPassword string
}

type Server struct {
	e               *echo.Echo
	config          *Configuration
	storage         *storage.Storage
	gameSessions    *gamesession.SessionStorage
	gameSessionsMu  sync.Mutex
	heartbeatPool   *heartbeatpool.Pool
	heartbeatTicker *time.Ticker
	heartbeatStop   chan bool
}

func New(config *Configuration, storage *storage.Storage) *Server {
	e := echo.New()
	e.Renderer = newPongo2Renderer(config)
	e.HTTPErrorHandler = httpErrorHandler
	e.Static(AssetsWebRoot, config.AssetsRoot)
	e.Static(UploadsWebRoot, config.UploadsRoot)
	e.Use(middleware.Recover())
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
		e:               e,
		config:          config,
		storage:         storage,
		gameSessions:    gamesession.NewSessionStorage(),
		heartbeatStop:   make(chan bool),
		heartbeatPool:   heartbeatpool.New(HeartbeatThreads),
		heartbeatTicker: time.NewTicker(HeartbeatInterval),
	}

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:    "form:_playtime_csrf",
		CookiePath:     "/",
		CookieName:     "_playtime_csrf",
		CookieHTTPOnly: true,
	}))
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "SAMEORIGIN",
		HSTSMaxAge:            0,
		HSTSExcludeSubdomains: false,
		HSTSPreloadEnabled:    false,
		ContentSecurityPolicy: "connect-src 'self' blob:; font-src 'self' data:; form-action 'self'; frame-ancestors 'self'; frame-src 'self'; img-src 'self' data:; manifest-src 'self'; media-src 'self'; object-src 'self'; child-src 'self' blob:",
	}))
	e.Use(s.contextCustomizationMiddleware)

	e.GET("/", s.index)

	e.HEAD(AssetsWebRoot+"*", s.assetsHead)
	e.HEAD(UploadsWebRoot+"*", s.uploadsHead)

	//authentication
	login := e.Group("/login")
	login.GET("", s.loginForm)
	login.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	login.POST("", s.loginSubmit)

	logout := e.Group("/logout")
	logout.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	logout.GET("", s.logout)

	//open netplay games
	open := e.Group("/open")
	open.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	open.GET("", s.open)

	//user profile
	profile := e.Group("/profile")
	profile.Use(s.authenticationRequiredMiddleware)
	profile.GET("", s.profileForm)
	profile.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	profile.POST("", s.profileSubmit)

	//settings
	settings := e.Group("/settings")
	settings.Use(s.authenticationRequiredMiddleware)
	settings.Use(s.settingsRequiredMiddleware)
	settings.GET("", s.settings)
	settings.GET("/general", s.settingsGeneralForm)
	settings.GET("/:platform", s.settingsByPlatformForm)
	settings.GET("/:platform/restore", s.settingsByPlatformRestoreDefaults)
	settings.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	settings.POST("/general", s.settingsGeneralSubmit)
	settings.POST("/:platform", s.settingsByPlatformSubmit)
	settings.POST("/:platform/restore", s.settingsByPlatformRestoreDefaultsSave)

	//users
	users := e.Group("/users")
	users.Use(s.authenticationRequiredMiddleware)
	users.Use(s.userControlAccessRequiredMiddleware)
	users.GET("", s.users)

	usersNew := users.Group("/new")
	usersNew.GET("", s.userNewForm)
	usersNew.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	usersNew.POST("", s.userNewSubmit)

	usersEdit := users.Group("/edit/:user_id")
	usersEdit.Use(s.userControlRequiredMiddleware)
	usersEdit.GET("", s.userEditForm)
	usersEdit.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	usersEdit.POST("", s.userEditSubmit)

	usersDelete := users.Group("/delete/:user_id")
	usersDelete.Use(s.userControlRequiredMiddleware)
	usersDelete.GET("", s.userDeleteForm)
	usersDelete.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	usersDelete.POST("", s.userDeleteSubmit)

	//games
	games := e.Group("/games")
	games.Use(s.authenticationRequiredMiddleware)
	games.GET("", s.games)

	gamesUpload := games.Group("/upload")
	gamesUpload.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1)))
	gamesUpload.POST("", s.gameUpload)

	gamesEmulationSettings := games.Group("/emulation-settings/:game_id")
	gamesEmulationSettings.Use(s.gameRequiredMiddleware)
	gamesEmulationSettings.GET("", s.gameEmulationSettingsForm)
	gamesEmulationSettings.GET("/restore", s.gameEmulationSettingsRestoreDefaults)
	gamesEmulationSettings.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	gamesEmulationSettings.POST("", s.gameEmulationSettingsSubmit)
	gamesEmulationSettings.POST("/restore", s.gameEmulationSettingsRestoreDefaultsSave)

	gamesEdit := games.Group("/edit/:game_id")
	gamesEdit.Use(s.gameRequiredMiddleware)
	gamesEdit.GET("", s.gameEditForm)
	gamesEdit.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	gamesEdit.POST("", s.gameEditSubmit)

	gamesDelete := games.Group("/delete/:game_id")
	gamesDelete.Use(s.gameRequiredMiddleware)
	gamesDelete.GET("", s.gameDeleteForm)
	gamesDelete.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	gamesDelete.POST("", s.gameDeleteSubmit)

	gamesNetplay := games.Group("/netplay/:game_id")
	gamesNetplay.Use(s.gameRequiredMiddleware)
	gamesNetplay.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	gamesNetplay.GET("/refresh-id", s.gameNetplayRefreshId)

	gamesControls := games.Group("/controls/:game_id")
	gamesControls.Use(s.gameRequiredMiddleware)
	gamesControls.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	gamesControls.POST("/save", s.gameControlsSave)

	uploadBatch := games.Group("/upload-batch/:upload_batch_id")
	uploadBatch.Use(s.uploadBatchRequiredMiddleware)
	uploadBatch.GET("", s.gameUploadBatchForm)
	uploadBatch.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	uploadBatch.POST("", s.gameUploadBatchSubmit)

	saveStates := games.Group("/save-states/:game_id")
	saveStates.Use(s.gameRequiredMiddleware)
	saveStates.GET("", s.saveStates)
	saveStates.GET("/list", s.saveStateList)

	saveStatesUpload := saveStates.Group("/upload")
	saveStatesUpload.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	saveStatesUpload.POST("", s.saveStateUpload)

	saveStateDelete := saveStates.Group("/delete/:save_state_id")
	saveStateDelete.Use(s.saveStateRequiredMiddleware)
	saveStateDelete.GET("", s.saveStateDeleteForm)
	saveStateDelete.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	saveStateDelete.POST("", s.saveStateDeleteSubmit)

	//play

	play := e.Group("/play/:game_id")
	play.Use(s.authenticationRequiredMiddleware)
	play.Use(s.gameRequiredMiddleware)
	play.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	play.GET("", s.play)

	//netplay

	netplay := e.Group("/netplay/:game_id/:netplay_session_id")
	netplay.Use(s.netplayGameRequiredMiddleware)
	netplay.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	netplay.GET("", s.netplay)
	netplay.GET("/ws", s.netplayWS)

	//misc

	version := e.Group("/version")
	version.Use(s.authenticationRequiredMiddleware)
	version.GET("", s.version)

	//netplay game session heartbeat

	go func() {
		for {
			select {
			case <-s.heartbeatTicker.C:
				s.netplayHeartbeat()
			case <-s.heartbeatStop:
				s.heartbeatTicker.Stop()
				s.heartbeatPool.Stop()
			}
		}
	}()

	return s
}

func (s *Server) Start() error {
	return s.e.Start(s.config.Listen)
}

func (s *Server) StopHeartbeat() error {
	s.heartbeatStop <- true
	return nil
}
