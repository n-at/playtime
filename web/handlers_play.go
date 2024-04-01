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

	emulatorSettings, err := s.getEmulatorSettings(context)
	if err != nil {
		return err
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
		for _, item := range storage.Platforms[game.Platform].Bios {
			if item.Name == emulatorSettings.Bios {
				bios = item
			}
		}
	}

	c.Response().Header().Add("Cross-Origin-Opener-Policy", "same-origin")
	c.Response().Header().Add("Cross-Origin-Embedder-Policy", "require-corp")

	netplayEnabled := s.config.NetplayEnabled && game.NetplayEnabled && len(game.NetplaySessionId) != 0

	if netplayEnabled {
		emulatorSettings.Volume = 1.0
		emulatorSettings.Buttons.Volume = false
		emulatorSettings.Buttons.Mute = false
	}

	return c.Render(http.StatusOK, "play", pongo2.Context{
		"_csrf_token":           c.Get("csrf"),
		"user":                  context.user,
		"game":                  game,
		"settings":              settings,
		"emulator_settings":     emulatorSettings,
		"bios":                  bios,
		"save_state":            saveState,
		"emulator_debug":        s.config.EmulatorDebug,
		"netplay_enabled":       netplayEnabled,
		"netplay_debug":         s.config.NetplayDebug,
		"netplay_turn_url":      s.config.TurnServerUrl,
		"netplay_turn_user":     s.config.TurnServerUser,
		"netplay_turn_password": s.config.TurnServerPassword,
	})
}
