package web

import (
	"errors"
	"fmt"
	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"playtime/storage"
)

func (s *Server) games(c echo.Context) error {
	context := c.(*PlaytimeContext)

	games, err := s.getGamesWithDataByUser(context.user)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "games", pongo2.Context{
		"_csrf_token":       c.Get("csrf"),
		"user":              context.user,
		"games_by_platform": s.groupGamesByPlatform(games),
		"netplay_enabled":   s.config.NetplayEnabled,
	})
}

func (s *Server) gameUpload(c echo.Context) error {
	context := c.(*PlaytimeContext)

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files, ok := form.File["games"]
	if !ok || len(files) == 0 {
		return errors.New("no games uploaded")
	}

	var uploadSize int64 = 0
	for _, file := range files {
		uploadSize += file.Size
	}

	if context.user.Quota > 0 && context.user.GetQuotaUsed()+uploadSize > context.user.Quota {
		return errors.New("disk quota exceeded")
	}

	var gameIds []string

	for _, file := range files {
		game := storage.Game{
			Id:                       storage.NewId(),
			UserId:                   context.user.Id,
			Name:                     file.Filename,
			OriginalFileName:         file.Filename,
			OriginalFileExtension:    getFileExtension(file.Filename),
			OriginalFileSize:         file.Size,
			Platform:                 "",
			OverrideEmulatorSettings: false,
			EmulatorSettings:         storage.DefaultEmulatorSettings(""),
		}

		if err := s.storage.SaveUploadedFile(file, game.Id, ""); err != nil {
			return err
		}

		if err := s.storage.GameUpload(game); err != nil {
			return err
		}

		gameIds = append(gameIds, game.Id)
	}

	uploadBatch := storage.UploadBatch{
		Id:      storage.NewId(),
		UserId:  context.user.Id,
		GameIds: gameIds,
	}

	if _, err := s.storage.UploadBatchSave(uploadBatch); err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, fmt.Sprintf("/games/upload-batch/%s", uploadBatch.Id))
}

func (s *Server) gameUploadBatchForm(c echo.Context) error {
	context := c.(*PlaytimeContext)

	games, err := s.storage.GameGetByUploadBatchId(context.uploadBatch.Id)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "upload_batch", pongo2.Context{
		"_csrf_token":  c.Get("csrf"),
		"user":         context.user,
		"upload_batch": context.uploadBatch,
		"games":        guessGameProperties(games),
		"platforms":    sortedPlatforms(),
	})
}

func (s *Server) gameUploadBatchSubmit(c echo.Context) error {
	context := c.(*PlaytimeContext)

	games, err := s.storage.GameGetByUploadBatchId(context.uploadBatch.Id)
	if err != nil {
		return err
	}

	for _, game := range games {
		game.Name = c.FormValue(fmt.Sprintf("name-%s", game.Id))
		game.Platform = c.FormValue(fmt.Sprintf("platform-%s", game.Id))
		game.EmulatorSettings = storage.DefaultEmulatorSettings(game.Platform)

		if _, err := s.storage.GameSave(game); err != nil {
			return err
		}
	}

	return c.Redirect(http.StatusFound, "/games")
}

func (s *Server) gameEditForm(c echo.Context) error {
	context := c.(*PlaytimeContext)

	return c.Render(http.StatusOK, "game_edit", pongo2.Context{
		"_csrf_token":     c.Get("csrf"),
		"user":            context.user,
		"game":            context.game,
		"platforms":       sortedPlatforms(),
		"netplay_enabled": s.config.NetplayEnabled,
	})
}

func (s *Server) gameEditSubmit(c echo.Context) error {
	context := c.(*PlaytimeContext)

	game := context.game
	game.Name = c.FormValue("name")
	game.OverrideEmulatorSettings = c.FormValue("override-settings") == "1"
	game.DisableCue = c.FormValue("disable-cue") == "1"
	game.NetplayEnabled = c.FormValue("netplay-enabled") == "1"

	newPlatform := c.FormValue("platform")
	if game.Platform != newPlatform {
		game.Platform = newPlatform
		game.EmulatorSettings = storage.DefaultEmulatorSettings(newPlatform)
	}

	if _, err := s.storage.GameSave(*game); err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, "/games")
}

func (s *Server) gameEmulationSettingsForm(c echo.Context) error {
	context := c.(*PlaytimeContext)

	if !context.game.OverrideEmulatorSettings {
		return errors.New("game does not override emulation settings")
	}

	platform := context.game.Platform

	return c.Render(http.StatusOK, "game_emulation_settings", pongo2.Context{
		"_csrf_token":          c.Get("csrf"),
		"user":                 context.user,
		"game":                 context.game,
		"settings":             context.game.EmulatorSettings,
		"shaders":              storage.Shaders,
		"platform":             storage.Platforms[platform],
		"bioses":               storage.Platforms[platform].Bios,
		"cores":                storage.Platforms[platform].Cores,
		"core_options":         storage.CoreOptionsByPlatform(platform),
		"fast_forward_ratios":  storage.FastForwardRatios,
		"slow_motion_ratios":   storage.SlowMotionRatios,
		"rewind_granularities": storage.RewindGranularities,
	})
}

func (s *Server) gameEmulationSettingsSubmit(c echo.Context) error {
	context := c.(*PlaytimeContext)
	game := context.game

	if !game.OverrideEmulatorSettings {
		return errors.New("game does not override emulation settings")
	}

	game.EmulatorSettings = settingsCollectFormData(c)

	if _, err := s.storage.GameSave(*game); err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, "/games")
}

func (s *Server) gameEmulationSettingsRestoreDefaults(c echo.Context) error {
	context := c.(*PlaytimeContext)
	game := context.game

	if !game.OverrideEmulatorSettings {
		return errors.New("game does not override emulation settings")
	}

	return c.Render(http.StatusOK, "game_emulation_settings_restore", pongo2.Context{
		"_csrf_token": c.Get("csrf"),
		"user":        context.user,
		"game":        context.game,
	})
}

func (s *Server) gameEmulationSettingsRestoreDefaultsSave(c echo.Context) error {
	context := c.(*PlaytimeContext)
	game := context.game

	if !game.OverrideEmulatorSettings {
		return errors.New("game does not override emulation settings")
	}

	game.EmulatorSettings = storage.DefaultEmulatorSettings(game.Platform)

	if _, err := s.storage.GameSave(*game); err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, "/games/emulation-settings/"+game.Id)
}

func (s *Server) gameDeleteForm(c echo.Context) error {
	context := c.(*PlaytimeContext)

	return c.Render(http.StatusOK, "game_delete", pongo2.Context{
		"_csrf_token": c.Get("csrf"),
		"user":        context.user,
		"game":        context.game,
	})
}

func (s *Server) gameDeleteSubmit(c echo.Context) error {
	context := c.(*PlaytimeContext)

	gameId := context.game.Id

	if err := s.storage.GameDeleteById(gameId); err != nil {
		return err
	}
	if err := s.storage.SaveStateDeleteByGameId(gameId); err != nil {
		log.Warnf("gameDeleteSubmit unable to delete deleted game %s save states: %s", gameId, err)
	}

	return c.Redirect(http.StatusFound, "/games")
}

func (s *Server) gameNetplayRefreshId(c echo.Context) error {
	context := c.(*PlaytimeContext)

	game := context.game
	prevSessionId := game.NetplaySessionId
	game.NetplaySessionId = storage.NewId()

	if _, err := s.storage.GameSave(*game); err != nil {
		return err
	}

	s.gameSessions.DeleteSession(prevSessionId)

	if c.QueryParam("return-to") == "play" {
		return c.Redirect(http.StatusFound, "/play/"+game.Id)
	} else {
		return c.Redirect(http.StatusFound, "/games")
	}
}

func (s *Server) gameControlsSave(c echo.Context) error {
	context := c.(*PlaytimeContext)
	game := context.game

	controls := settingsCollectControls(c)

	if game.OverrideEmulatorSettings {
		game.EmulatorSettings.Controls = controls
		if _, err := s.storage.GameSave(*game); err != nil {
			return err
		}
	} else {
		settings, err := s.storage.SettingsGetByUserId(context.session.UserId)
		if err != nil {
			return err
		}
		platformSettings, ok := settings.EmulatorSettings[game.Platform]
		if !ok {
			return errors.New("platform not found")
		}

		platformSettings.Controls = controls
		settings.EmulatorSettings[game.Platform] = platformSettings
		if _, err := s.storage.SettingsSave(settings); err != nil {
			return err
		}
	}

	return c.JSON(http.StatusOK, nil)
}
