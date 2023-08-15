package web

import (
	"errors"
	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	"net/http"
	"playtime/storage"
)

func (s *Server) play(c echo.Context) error {
	context := c.(*PlaytimeContext)

	game, err := s.getGameWithDataById(context.user, context.game.Id)
	if err != nil {
		return err
	}

	if len(game.Platform) == 0 {
		return errors.New("game platform is undefined")
	}

	settings, err := s.storage.SettingsGetByUserId(context.user.Id)
	if err != nil {
		return err
	}

	emulatorSettings := settings.EmulatorSettings[game.Platform]
	if game.OverrideEmulatorSettings {
		emulatorSettings = game.EmulatorSettings
	}

	saveState := storage.SaveStateWithData{}
	saveStateId := c.QueryParam("state")
	if len(saveStateId) != 0 {
		saveState, err = s.getSaveStateWithDataById(context.user, saveStateId)
		if err != nil {
			return err
		}
		if saveState.GameId != game.Id {
			return errors.New("save state belongs to different game")
		}
		if saveState.UserId != context.user.Id {
			return errors.New("save state belongs to different user")
		}
	}

	bios := storage.Bios{}
	if len(emulatorSettings.Bios) != 0 {
		for _, item := range storage.Bioses[game.Platform] {
			if item.Name == emulatorSettings.Bios {
				bios = item
			}
		}
	}

	return c.Render(http.StatusOK, "play", pongo2.Context{
		"user":                  context.user,
		"game":                  game,
		"settings":              settings,
		"emulator_settings":     emulatorSettings,
		"bios":                  bios,
		"save_state":            saveState,
		"netplay_enabled":       s.config.NetplayEnabled && game.NetplayEnabled && len(game.NetplaySessionId) != 0,
		"netplay_turn_url":      s.config.TurnServerUrl,
		"netplay_turn_user":     s.config.TurnServerUser,
		"netplay_turn_password": s.config.TurnServerPassword,
	})
}
